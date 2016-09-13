package isitrandom

import (
	"io"
	"testing"
)

type test interface {
	next(b byte)
	p() float64
}

func buildTest(t test) func(c chan byte) float64 {
	return func(c chan byte) float64 {
		for b := range c {
			t.next(b)
		}
		return t.p()
	}
}

func buildIOReaderTest(t test) func(r io.Reader) float64 {
	return func(r io.Reader) float64 {
		p := make([]byte, 1)
		var err error
		for err != nil {
			_, err = r.Read(p)
			t.next(p[0])
		}
		return t.p()
	}
}

var tests = []struct {
	name string
	t    test
}{
	// {"frequency", &frequency{}},
	{"serial", &serial{}},
}

// TestRNG tests all the tests available and fails if any p-value is too low.
// This runs all tests in paraellel. That might not be good for reading.........
func TestRNG(t *testing.T, rng io.Reader) {
	c := readerToChannel(rng)
	s := split(c, len(tests))
	for i, a := range tests {
		rngt := buildTest(a.t)
		t.Run(a.name, func(t *testing.T) {
			p := rngt(s[i])
			if p < 0.05 {
				t.Errorf("Fail %s: expected p-value > 0.05, got %f", a.name, p)
			}
		})
	}
}

type channelReader struct {
	r io.Reader
	c chan byte
}

func (cr *channelReader) read() {
	var err error
	p := make([]byte, 1)
	for {
		_, err = cr.r.Read(p)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			close(cr.c)
			break
		}
		cr.c <- p[0]
	}
}

func readerToChannel(r io.Reader) chan byte {
	cr := channelReader{r, make(chan byte)}
	go cr.read()
	return cr.c
}

func split(input chan byte, n int) (output []chan byte) {
	for i := 0; i < n; i++ {
		output = append(output, make(chan byte))
	}
	go func() {
		for b := range input {
			for i := 0; i < n; i++ {
				output[i] <- b // this will block until someone recieves?
			}
		}
		for i := 0; i < n; i++ {
			close(output[i])
		}
	}()
	return output
}
