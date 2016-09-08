package main

import (
	"io"
	"testing"
)

// Test is a function which takes a random number generator and attempts
// to perform a randomness test. It should never fail. It should return a
// p-value indictaing whether it passed the test.
type Test func(rng io.Reader) float64

var tests = []struct {
	name string
	test Test
}{
	{"frequency", FrequencyTest},
}

// TestRNG tests all the tests available and fails if any p-value is too low.
// This runs all tests in paraellel. That might not be good for reading.........
func TestRNG(t *testing.T, rng io.Reader) {
	for _, a := range tests {
		t.Run(a.name, func(t *testing.T) {
			p := a.test(rng)
			if p > 0.05 {
				t.Errorf("Fail %s: expected p-value ≤ 0.05, got %f", a.name, p)
			}
		})
	}
}
