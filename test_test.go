package main

import "testing"

type alternatingRNG struct {
	x byte
}

func (rng alternatingRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		var b byte
		for j := 0; j < 8; j++ {
			b = b<<1 | rng.x
			rng.x = 1 - rng.x
		}
		p[i] = b
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
