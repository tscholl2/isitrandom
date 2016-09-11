package isitrandom

import "testing"

func TestSerialTest(t *testing.T) {
	/*
		t.Run("correct on small input", func(t *testing.T) {
			var buf = bytes.NewBuffer([]byte{0x00})
			_, statistic := SerialTestN(buf, 1)
			if statistic != 64.0/8 {
				t.Errorf("got: %f explected: %f", statistic, 64.0/8)
			}
		})
		t.Run("reads whole ", func(t *testing.T) {
			var buf = bytes.NewBuffer([]byte{0x00, 0xff})
			_, statistic := SerialTestN(buf, -1)
			if statistic != 0 {
				t.Errorf("got: %f explected: %d", statistic, 0)
			}
		})
		t.Run("panics on not enough values", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("expected a panic, but didnt happen")
				}
			}()
			var buf = bytes.NewBuffer([]byte{0x00, 0xff})
			SerialTestN(buf, 3)
		})
	*/
}
