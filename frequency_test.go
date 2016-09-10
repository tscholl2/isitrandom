package isitrandom

import (
	"bytes"
	"fmt"
	"testing"
)

type alternatingRNG struct{}

func (rng alternatingRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xaa // 10101010
	}
	return len(p), nil
}

type slightlyAlternatingRNG struct{}

func (rng slightlyAlternatingRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xbb // 10111011
	}
	return len(p), nil
}

type constantRNG struct{}

func (rng constantRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xff // 11111111
	}
	return len(p), nil
}

func TestFrequencyTest(t *testing.T) {
	var p, statistic float64
	var targetP, targetStatistic float64

	for _, N := range []int{2, 4, 8, 100} {
		t.Run(fmt.Sprintf("alternating_N=%d", N), func(t *testing.T) {
			p = FrequencyTest(alternatingRNG{})
			targetP = 0.0
			if p != 0.0 {
				t.Errorf("alternatingRNG, Expected p-value of %f, got %f", targetP, p)
			}
		})
	}

	p = FrequencyTest(slightlyAlternatingRNG{})
	targetP = 0.9999
	if p != 0.9999 {
		t.Errorf("slightlyAlternatingRNG, Expected p-value of %f, got %f", targetP, p)
	}

	p = FrequencyTest(constantRNG{})
	targetP = 0.9999
	if p != 0.9999 {
		t.Errorf("constantRNG, Expected p-value of %f, got %f", targetP, p)
	}

	// 1110001100010001010011101111001001001001 x 4
	menezesRNG := bytes.NewBuffer([]byte{0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49})
	p, statistic = FrequencyTestN(menezesRNG, menezesRNG.Len())
	targetP = 0.180000
	if p-targetP > MACHINE_EPSILON {
		t.Errorf("menezesRNG, Expected p-value of %f, got %f", targetP, p)
	}
	targetStatistic = 0.4
	if statistic-targetStatistic > MACHINE_EPSILON {
		t.Errorf("menezesRNG, Expected targetStatistic of %f, got %f", targetStatistic, statistic)
	}

}