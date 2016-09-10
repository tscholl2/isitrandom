package isitrandom

import (
	"bytes"
	"io"
)

// SerialTest tests whether the number of occurences of
// 00, 01, 10, and 11 as subsequences are approximately the same
// as would be expected for a random sequence.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181
func SerialTest(rng io.Reader) float64 {
	N := 100
	b := make([]byte, N)
	for i := 0; i < N; i++ {
		rng.Read(b)
	}
	pValue, _ := SerialTestN(bytes.NewBuffer(b))
	return pValue
}

// SerialTestN tests the number of 2-bit subsequences
// and compares it to the expected number. It will read N
// *bytes*. If N = -1, it will read until EOF.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181.
func SerialTestN(rng io.Reader) (float64, float64) {
	// counts           00 01 10 11
	counts := []float64{0, 0, 0, 0}
	started := false
	lastBit := false
	buf := new(bytes.Buffer)
	buf.ReadFrom(rng)
	data := buf.Bytes()
	r := New(bytes.NewBuffer(data))
	for i := 0; i < len(data); i++ {
		for j := 0; j < 8; j++ {
			bit, _ := r.ReadBit()
			// if bit {
			// 	fmt.Printf("1")
			// } else {
			// 	fmt.Printf("0")
			// }
			if !started {
				started = true
				lastBit = bit
				continue
			}
			if !lastBit && !bit { // 00
				counts[0]++
			} else if !lastBit && bit { // 01
				counts[1]++
			} else if lastBit && !bit { // 10
				counts[2]++
			} else if lastBit && bit { // 11
				counts[3]++
			}
			lastBit = bit
		}
	}

	n := counts[0] + counts[1] + counts[2] + counts[3] + 1
	n0 := counts[0] + counts[1]
	n1 := counts[2] + counts[3] + 1
	statistic := 4/(n-1)*(counts[0]*counts[0]+counts[1]*counts[1]+counts[2]*counts[2]+counts[3]*counts[3]) - (2/n)*(n0*n0+n1*n1) + 1
	return chisquared(statistic, 2), statistic
}

// Following code is from http://stackoverflow.com/questions/29583024/reading-8-bits-from-a-reader-in-golang?answertab=votes#tab-top
type BitReader struct {
	reader io.ByteReader
	byte   byte
	offset byte
}

func New(r io.ByteReader) *BitReader {
	return &BitReader{r, 0, 0}
}

func (r *BitReader) ReadBit() (bool, error) {
	if r.offset == 8 {
		r.offset = 0
	}
	if r.offset == 0 {
		var err error
		if r.byte, err = r.reader.ReadByte(); err != nil {
			return false, err
		}
	}
	bit := (r.byte & (0x80 >> r.offset)) != 0
	r.offset++
	return bit, nil
}

func (r *BitReader) ReadUint(nbits int) (uint64, error) {
	var result uint64
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.ReadBit()
		if err != nil {
			return 0, err
		}
		if bit {
			result |= 1 << uint(i)
		}
	}
	return result, nil
}
