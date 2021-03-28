package tsparser

import (
	"bufio"
	"bytes"
	"os"
)

type TsPacket []byte

func Scan(path string, pid int, tidRange TableIdRange) [][]byte {
	fp, _ := os.Open(path)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Split(splitPacket)

	var sectionBytes []byte
	var sectionBytesList [][]byte

	var sectionLength int = -1
	var payload []byte
	var tid byte
	var pointer int

	for scanner.Scan() {
		var packet TsPacket = scanner.Bytes()
		header := ParseTsHeader(packet)

		switch {
		case header.PID != pid:
			continue
		case !header.HasPayload():
			continue
		}

		payload = packet.getPayload(header.PayloadUnitStart)

		// Append remaining data
		if len(sectionBytes) > 0 {
			if sectionLength >= len(payload) {
				// No new start
				sectionBytes = append(sectionBytes, payload...)
				sectionLength -= len(payload)
				continue
			} else {
				// Has new start
				if header.PayloadUnitStart {
					payload = payload[1:] // remove pointer field
				}
				sectionBytes = append(sectionBytes, payload[:sectionLength]...)
				sectionBytesList = append(sectionBytesList, sectionBytes)
				sectionBytes = []byte{}
			}
		}

		if !header.PayloadUnitStart {
			continue
		}

		payload = packet.getPayload(header.PayloadUnitStart)

		pointer = packet.getPointerField(header.PayloadUnitStart)
		payload = payload[pointer+1:]

		tid = payload[0]
		if tid < tidRange.Start || tid > tidRange.End {
			continue
		}

		sectionLength = int(payload[2]) | int(payload[1]&0x0F)<<8 + 3
		if sectionLength <= len(payload) {
			sectionBytes = append(sectionBytes, payload[:sectionLength]...)
			sectionBytesList = append(sectionBytesList, sectionBytes)
			sectionBytes = []byte{}
			continue
		}

		sectionBytes = append(sectionBytes, payload...)
		sectionLength -= len(payload)
	}

	return sectionBytesList
}

func splitPacket(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if len(data) < tsPacketSize {
		return 0, nil, nil
	}
	i := bytes.IndexByte(data[tsPacketSize:], byte(0x47))
	if i >= 0 {
		return i + tsPacketSize, data[0 : i+tsPacketSize], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func (p TsPacket) getPointerField(payloadUnitStart bool) int {
	if !payloadUnitStart {
		return -1
	}
	return int(p[4])
}

func (p TsPacket) getPayload(payloadUnitStart bool) []byte {
	return p[4:]
}
