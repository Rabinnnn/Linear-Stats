package maths

// calculate linear regression
func LinearRegression(x, y []float64) (float64, float64) {
	n := float64(len(x)) // number of points (x axis)
	sumx := SumXaxis(x) // sum of all values on x axis
	sumy := SumYaxis(y) // sum of all values on y axis
	sumxy := XYSumProduct(x, y)  // Sum of products of corresponding x and y values
	sumxsq := SumXSquares(x) // Sum of the squares of all x values

	b := (n*sumxy - sumx*sumy) / (n*sumxsq - sumx*sumx)
	a := (sumy - b*sumx) / n
	return b, a
}

// calculate sum of all values on y axis
func SumYaxis(y []float64) float64 {
	temp := y[0]

	for i := 1; i < len(y); i++ {
		temp += y[i]
	}
	return temp
}

// calculate sum of all values on x axis
func SumXaxis(x []float64) float64 {
	temp := x[0]

	for i := 1; i < len(x); i++ {
		temp += x[i]
	}
	return temp
}

// calculate sum of the squares of all x values
func SumXSquares(x []float64) float64 {
	temp := 0.0

	for i := 0; i < len(x); i++ {
		temp += x[i] * x[i]
	}
	return temp
}

// calculate sum of products of corresponding x and y values
func XYSumProduct(x, y []float64) float64 {
	temp := 0.0

	for i := 0; i < len(x); i++ {
		temp += x[i] * y[i]
	}
	return temp
}
