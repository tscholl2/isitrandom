package isitrandom

import "testing"

func TestSerialTest(t *testing.T) {
	t.Run("correct on small input", func(t *testing.T) {
		ft := &serial{}
		short := []byte{
			0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3, 0x11, 0x4e,
			0xf2, 0x49, 0xe3, 0x11, 0x4e, 0xf2, 0x49, 0xe3,
			0x11, 0x4e, 0xf2, 0x49}
		for _, val := range short {
			ft.next(val)
		}
		statistic := ft.statistic()
		if statistic-0.625157 > 0.00001 {
			t.Errorf("got: %f explected: %f", statistic, 0.625157)
		}
		pValues := ft.p()
		if pValues-0.9 > 0.00001 {
			t.Errorf("got: %f explected: %f", pValues, 0.9)
		}
	})
}
