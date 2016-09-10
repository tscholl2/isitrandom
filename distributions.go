package main

func normZeroOne(x float64) (y float64) {
	for i, val := range NormalTable.X {
		if x > val {
			return NormalTable.P[i]
		}
	}
	return -1
}

func chisquared(x float64, df int64) (y float64) {
	for i, val := range ChiSquareTable.X[df] {
		if x > val {
			return ChiSquareTable.P[i]
		}
	}
	return -1.0
}
