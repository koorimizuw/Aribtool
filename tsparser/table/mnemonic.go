package table

import (
	"time"

	"github.com/koorimizuw/B24Decoder/b24decoder"
)

type Mnemonic []byte

func (s Section) uimsbf(index, length int) Mnemonic {
	var m Mnemonic
	for length > 0 {
		start := index / 8
		pos := index % 8
		rem := 8 - pos
		if length <= rem {
			m = append(m, s[start]<<pos>>(8-length))
			break
		}
		m = append(m, s[start]<<pos>>pos)
		index += rem
		length -= rem
	}
	return m
}

// 処理はuimsbfと一緒？
func (s Section) bslbf(index, length int) Mnemonic {
	var m Mnemonic
	for length > 0 {
		start := index / 8
		pos := index % 8
		rem := 8 - pos
		if length <= rem {
			m = append(m, s[start]<<pos>>(8-length))
			break
		}
		m = append(m, s[start]<<pos>>pos)
		index += rem
		length -= rem
	}
	return m
}

func (m Mnemonic) toByte() byte {
	return m[0]
}

func (m Mnemonic) toBool() bool {
	return m[0] == 1
}

func (m Mnemonic) toNumber() int {
	num := 0
	for i, v := range m {
		num += int(v) << ((len(m) - i - 1) * 8)
	}
	return num
}

func (m Mnemonic) ToString() string {
	return b24decoder.Decode(m)
}

func (m Mnemonic) toTime() time.Time {
	if len(m) < 5 {
		panic("Too few bytes")
	}
	year, month, day := decodeMJD(m[0], m[1])
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, JST)
	return t.Add(bcd(m[2], m[3], m[4]))
}

func (m Mnemonic) toDuration() int {
	if len(m) < 3 {
		panic("Too few bytes")
	}
	return int(bcd(m[0], m[1], m[2]).Minutes())
}

func (m Mnemonic) toLanguageCode() string {
	if len(m) < 3 {
		panic("Too few bytes")
	}
	return string(m[0]) + string(m[1]) + string(m[2])
}
