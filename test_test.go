package isitrandom

import "testing"

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

	p = FrequencyTest(alternatingRNG{})
	if p != 0.0 {
		t.Errorf("Expected %f, got %f", 0.0, p)
	}

	p = FrequencyTest(slightlyAlternatingRNG{})
	if p != 0.5 {
		t.Errorf("slightlyAlternatingRNG, Expected %f, got %f", 0.9999, p)
	}

	p = FrequencyTest(constantRNG{})
	if p != 0.9999 {
		t.Errorf("Expected %f, got %f", 0.9999, p)
	}
}
