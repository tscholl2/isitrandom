package examples

import (
	"testing"

	"github.com/tscholl2/isitrandom"
)

type constantRNG struct{}

func (constantRNG) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0xff // 11111111
	}
	return len(p), nil
}

func TestConstantRNG(t *testing.T) {
	isitrandom.TestRNG(t, constantRNG{})
}
