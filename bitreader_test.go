package isitrandom

import (
	"bytes"
	"testing"
)

func TestBitReader(t *testing.T) {
	t.Run("test reading bits one at a time", func(t *testing.T) {
		data := []byte{0xe3, 0x11, 0x4e, 0xf2, 0x49}
		bitStringExample := "1110001100010001010011101111001001001001"
		bitString := ""
		r := New(bytes.NewBuffer(data))
		for i := 0; i < len(data); i++ {
			for j := 0; j < 8; j++ {
				bit, _ := r.ReadBit()
				if bit {
					bitString += "1"
				} else {
					bitString += "0"
				}
			}
		}
		if bitString != bitStringExample {
			t.Errorf("got: %s expected: %s", bitString, bitStringExample)
		}
	})
}
