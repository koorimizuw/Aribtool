package table

import (
	"fmt"
	"math"
	"time"
)

var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

// go-arib/arib/eit.go
func decodeMJD(b1, b2 byte) (int, int, int) {
	mjd := float64(uint16(b1&0xFF)<<8 | uint16(b2&0xFF))
	y := math.Trunc((mjd - 15078.2) / 365.25)
	m := math.Trunc((mjd - 14956.1 - math.Trunc(y*365.25)) / 30.6001)
	d := mjd - 14956 - math.Trunc(y*365.25) - math.Trunc(m*30.6001)
	var k float64
	if m == 14 || m == 15 {
		k = 1
	}
	y += k
	m = m - 1 - k*12
	return 1900 + int(y), int(m), int(d)
}

// go-arib/arib/eit.go
func bcd(b1, b2, b3 byte) time.Duration {
	hms := fmt.Sprintf("%02Xh%02Xm%02Xs", b1, b2, b3)
	d, err := time.ParseDuration(hms)
	if err != nil {
		return time.Duration(-1)
	}
	return d
}
