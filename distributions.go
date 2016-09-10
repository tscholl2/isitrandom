package isitrandom

func normZeroOne(x float64) (y float64) {
	for i, val := range NormalTable.X {
		if x > val {
			return 1.0 - NormalTable.P[i]
		}
	}
	return 0.0
}

func chisquared(x float64, df int64) (y float64) {
	for i, val := range ChiSquareTable.X[df] {
		if x > val {
			return 1.0 - ChiSquareTable.P[i]
		}
	}
	return 0.0
}
