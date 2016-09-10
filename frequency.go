package isitrandom

import (
	"fmt"
	"io"
)

// FrequencyTest tests the number of 1s and 0s it sees
// and compares it to the expected number.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181
func FrequencyTest(rng io.Reader) float64 {
	return FrequencyTestN(rng, 10000)
}

// FrequencyTestN tests the number of 1s and 0s it sees
// and compares it to the expected number. It will read N
// *bytes*. If N = -1, it will read until EOF.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181.
func FrequencyTestN(rng io.Reader, N int) float64 {
	ones := float64(0)
	b := make([]byte, 1)
	bytesRead := 0
	var err error
	for bytesRead < N || (N < 0 && err != io.EOF) {
		if _, err := rng.Read(b); err != nil && err != io.EOF {
			panic(err)
		} else if N > -1 && err == io.EOF && bytesRead < N-1 {
			panic(fmt.Errorf("not enough bytes: read %d, want: %d", bytesRead, N))
		}
		bytesRead++
		for b[0] > 0 {
			if b[0]&1 == 1 {
				ones++
			}
			b[0] = b[0] >> 1
		}
	}
	n := float64(8 * bytesRead)
	zeros := n - ones
	return chisquared((zeros-ones)*(zeros-ones)/(ones+zeros), 1)
}
