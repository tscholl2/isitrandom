package main

type normalLookupTable struct {
	P []float64
	X []float64
}

var (
	// NormalTable stores the frequently needed values for looking up N(0,1) tests
	NormalTable = normalLookupTable{
		[]float64{0.000100, 0.001000, 0.002000, 0.005000, 0.010000, 0.020000, 0.025000, 0.050000, 0.100000, 0.200000, 0.975000, 0.995000, 0.999000, 0.999900},
		[]float64{-3.719016, -3.090232, -2.878162, -2.575829, -2.326348, -2.053749, -1.959964, -1.644854, -1.281552, -0.841621, 1.959964, 2.575829, 3.090232, 3.719016},
	}
)
