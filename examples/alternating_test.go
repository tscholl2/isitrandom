package examples

import (
	"testing"

	iir "github.com/tscholl2/is-it-random"
)

type alternatingRNG struct {
	x byte
}

func (rng alternatingRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xaa // 10101010
	}
	return len(p), nil
}

func TestAlternatingRNG(t *testing.T) {
	iir.TestRNG(t, alternatingRNG{})
}
