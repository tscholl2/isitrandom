package isitrandom

import "io"

type frequency struct {
	ones  int64
	bytes int64
}

func (f *frequency) next(b byte) {
	f.bytes++
	for b > 0 {
		f.ones += int64(b & 1)
		b = b >> 1
	}
}

func (f *frequency) statistic() float64 {
	zeros := 8*f.bytes - f.ones
	return float64((zeros - f.ones)) / float64(f.ones+zeros) * float64(zeros-f.ones)
}

func (f *frequency) p() float64 {
	s := f.statistic()
	return chisquared(s, 1)
}

// FrequencyTest tests the number of 1s and 0s it sees
// and compares it to the expected number. It will read N
// *bytes*. If N = -1, it will read until EOF.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181.
func FrequencyTest(rng io.Reader) float64 {
	return buildIOReaderTest(&frequency{})(rng)
}
