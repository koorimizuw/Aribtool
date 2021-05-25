package tsparser

type TsHeader struct {
	SyncByte                   byte
	TransportError             bool
	PayloadUnitStart           bool
	TransportPriority          byte
	PID                        int
	TransportScramblingControl byte
	AdaptationFieldControl     byte
	ContinuityCounter          int
}

func ParseTsHeader(b []byte) TsHeader {
	return TsHeader{
		SyncByte:                   b[0],
		TransportError:             b[1]&0x80>>7 == 1,
		PayloadUnitStart:           b[1]&0x40>>6 == 1,
		TransportPriority:          b[1] & 0x20 >> 5,
		PID:                        int(b[2]) | (int(b[1])&0x1f)<<8,
		TransportScramblingControl: b[3] & 0xc0 >> 6,
		AdaptationFieldControl:     b[3] & 0x30 >> 4,
		ContinuityCounter:          int(b[3] & 0x0F),
	}
}

var PidMap = map[string]int{
	ProgramAssociationSection: 0x0,
	EventSection:              0x12,
	CurrentEventSection:       0x12,
	ScheduleEventSection:      0x12,
	// ...
}

func (h *TsHeader) HasAdaptationField() bool {
	return h.AdaptationFieldControl == 0x02 || h.AdaptationFieldControl == 0x03
}

func (h *TsHeader) HasPayload() bool {
	return h.AdaptationFieldControl == 0x01 || h.AdaptationFieldControl == 0x03
}

func (h *TsHeader) PrevContinuityCounter() int {
	if h.ContinuityCounter == 0 {
		return 15
	}
	return h.ContinuityCounter - 1
}
