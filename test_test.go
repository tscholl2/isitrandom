package isitrandom

import (
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
	var p float64
	var targetP float64
	for _, N := range []int{2, 4, 8, 100} {
		t.Run(fmt.Sprintf("alternating_N=%d", N), func(t *testing.T) {
			p = FrequencyTest(alternatingRNG{})
			targetP = 0.0
			if p != 0.0 {
				t.Errorf("alternatingRNG, Expected %f, got %f", targetP, p)
			}
		})
	}

	p = FrequencyTest(slightlyAlternatingRNG{})
	targetP = 0.86
	if p != 0.86 {
		t.Errorf("slightlyAlternatingRNG, Expected %f, got %f", targetP, p)
	}

	p = FrequencyTest(constantRNG{})
	targetP = 0.999
	if p != 0.999 {
		t.Errorf("constantRNG, Expected %f, got %f", targetP, p)
	}

}
