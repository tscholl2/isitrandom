package isitrandom

import "testing"

type alternatingRNG struct{}

func (rng alternatingRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xaa // 10101010
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
	// TODO fix this
	if p != 0.5 {
		t.Errorf("Expected %f, got %f", 0.5, p)
	}
	// TODO fix this
	p = FrequencyTest(constantRNG{})
	if p < 0.5 {
		t.Errorf("Expected %f, got %f", 0.5, p)
	}
}
