package isitrandom

import (
	"bytes"
	"testing"
)

func TestSerialTest(t *testing.T) {
	var p, statistic float64
	var targetP, targetStatistic float64

	// 1110001100010001010011101111001001001001 x 4
	menezesRNG := bytes.NewBuffer([]byte{0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49})
	p, statistic = SerialTestN(menezesRNG, menezesRNG.Len())
	targetP = 0.180000
	if p != targetP {
		t.Errorf("menezesRNG, Expected p-value of %f, got %f", targetP, p)
	}
	targetStatistic = 0.6252
	if statistic != targetStatistic {
		t.Errorf("menezesRNG, Expected targetStatistic of %f, got %f", targetStatistic, statistic)
	}

}
