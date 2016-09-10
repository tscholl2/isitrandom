package main

import "testing"

type alternatingRNG struct {
	x byte
}

func (rng alternatingRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xaa // 10101010
	}
	return len(p), nil
}

func TestFrequencyTest(t *testing.T) {
	p := FrequencyTest(alternatingRNG{})
	// TODO fix this
	if p != 0.5 {
		t.Errorf("Expected %f, got %f", 0.5, p)
	}
}

func TestTestRNG(t *testing.T) {
	TestRNG(t, alternatingRNG{})
}
