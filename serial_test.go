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
	p, statistic = SerialTestN(menezesRNG)
	targetP = 0.9
	if p-targetP > MACHINE_EPSILON {
		t.Errorf("menezesRNG, Expected p-value of %f, got %f", targetP, p)
	}
	targetStatistic = 0.62515
	if statistic-targetStatistic > MACHINE_EPSILON {
		t.Errorf("menezesRNG, Expected targetStatistic of %f, got %f", targetStatistic, statistic)
	}

	p, statistic = SerialTestN(randomRNG)
	targetP = 1.0
	if p-targetP > MACHINE_EPSILON {
		t.Errorf("randomRNG, Expected p-value of %f, got %f", targetP, p)
	}
	targetStatistic = 0.62515
	if statistic-targetStatistic > MACHINE_EPSILON {
		t.Errorf("randomRNG, Expected targetStatistic of %f, got %f", targetStatistic, statistic)
	}

	p = SerialTest(slightlyAlternatingRNG{})
	targetP = 0.0001
	if p-targetP > MACHINE_EPSILON {
		t.Errorf("slightlyAlternatingRNG, Expected p-value of %f, got %f", targetP, p)
	}

	p = SerialTest(constantRNG{})
	targetP = 0.0001
	if p-targetP > MACHINE_EPSILON {
		t.Errorf("constantRNG, Expected p-value of %f, got %f", targetP, p)
	}

	p = SerialTest(alternatingRNG{})
	targetP = 0.0001
	if p-targetP > MACHINE_EPSILON {
		t.Errorf("alternatingRNG, Expected p-value of %f, got %f", targetP, p)
	}
}
