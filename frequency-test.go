package isitrandom

import (
	"fmt"
	"io"
)

const frequencyTestN = 10000

// FrequencyTest tests the number of 1s and 0s it sees
// and compares it to the expected number.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181
func FrequencyTest(rng io.Reader) float64 {
	ones := float64(0)
	b := make([]byte, 1)
	for i := 0; i < frequencyTestN; i++ {
		if _, err := rng.Read(b); err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF && i < frequencyTestN-1 {
			panic(fmt.Errorf("not enough bytes: read %d, want: %d", i, frequencyTestN))
		}
		for b[0] > 0 {
			if b[0]&1 == 1 {
				ones++
			}
			b[0] = b[0] >> 1
		}
	}
	n := float64(8 * frequencyTestN)
	zeros := n - ones
	return chisquared((zeros-ones)*(zeros-ones)/n, 1)
}
