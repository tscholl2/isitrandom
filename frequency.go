package isitrandom

import (
	"fmt"
	"io"
)

// FrequencyTest tests the number of 1s and 0s it sees
// and compares it to the expected number.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181
func FrequencyTest(rng io.Reader) float64 {
	pValue, _ := FrequencyTestN(rng, 10000)
	return pValue
}

// FrequencyTestN tests the number of 1s and 0s it sees
// and compares it to the expected number. It will read N
// *bytes*. If N = -1, it will read until EOF.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181.
func FrequencyTestN(rng io.Reader, N int) (float64, float64) {
	ones, zeros, n := float64(0), float64(0), float64(0)
	var m int
	b := make([]byte, 1)
	bytesRead := 0
	var err error
	for bytesRead < N || (N < 0 && err != io.EOF) {
		m, err = rng.Read(b)
		bytesRead += m
		// if there is a weird error, panic
		if err != nil && err != io.EOF {
			panic(err)
		}
		// if we get to the end before we read all the bytes, panic
		if N > -1 && err == io.EOF && bytesRead < N {
			panic(fmt.Errorf("not enough bytes: read %d, want: %d", bytesRead, N))
		}
		// if we are at the end, were done
		if err == io.EOF {
			break
		}
		for j := 0; j < 8; j++ {
			if b[0]&1 == 1 {
				ones++
			} else {
				zeros++
			}
			b[0] = b[0] >> 1
			n++
		}
	}
	return chisquared((zeros-ones)*(zeros-ones)/(ones+zeros), 1), (zeros - ones) * (zeros - ones) / (ones + zeros)
}
