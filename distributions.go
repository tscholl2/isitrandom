package main

func chisquared(x float64, df int64) (y float64) {
	for i, val := range ChiSquareTable.X[df] {
		if x > val {
			return ChiSquareTable.P[i]
		}
	}
	return -1.0
}
