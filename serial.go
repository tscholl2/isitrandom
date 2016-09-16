package isitrandom

import (
	"bytes"
	"io"
)

type serial struct {
	n00, n01, n10, n11 float64
	started            bool
	lastBit            bool
}

func (f *serial) next(b byte) {
	r := New(bytes.NewBuffer([]byte{b}))
	for j := 0; j < 8; j++ {
		bit, _ := r.ReadBit()
		// if bit {
		// 	fmt.Printf("1")
		// } else {
		// 	fmt.Printf("0")
		// }
		if !f.started {
			f.started = true
			f.lastBit = bit
			continue
		}
		if !f.lastBit && !bit { // 00
			f.n00++
		} else if !f.lastBit && bit { // 01
			f.n01++
		} else if f.lastBit && !bit { // 10
			f.n10++
		} else if f.lastBit && bit { // 11
			f.n11++
		}
		f.lastBit = bit
	}
}

func (f *serial) statistic() float64 {
	n := f.n00 + f.n01 + f.n10 + f.n11 + 1
	n0 := f.n00 + f.n01
	n1 := f.n10 + f.n11 + 1
	statistic := 4.0/(n-1)*(f.n00*f.n00+f.n01*f.n01+f.n10*f.n10+f.n11*f.n11) - (2/n)*(n0*n0+n1*n1) + 1
	return statistic
}

func (f *serial) p() float64 {
	s := f.statistic()
	return chisquared(s, 2)
}

// SerialTest tests whether the number of occurences of
// 00, 01, 10, and 11 as subsequences are approximately the same
// as would be expected for a random sequence.
// See http://cacr.uwaterloo.ca/hac/about/chap5.pdf pg 181
func SerialTest(rng io.Reader) float64 {
	return buildIOReaderTest(&serial{})(rng)
}

func SerialP(short []byte) float64 {
	ft := &serial{}
	for _, val := range short {
		ft.next(val)
	}
	return ft.p()
}
