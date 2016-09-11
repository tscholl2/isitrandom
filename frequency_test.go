package isitrandom

import "testing"

func TestFreqency(t *testing.T) {
	t.Run("correct on small input", func(t *testing.T) {
		ft := &frequency{}
		ft.next(0x00)
		statistic := ft.statistic()
		if statistic != 64.0/8 {
			t.Errorf("got: %f explected: %f", statistic, 64.0/8)
		}
	})
}
