package maths

import "fmt"

func Average(sum float64, length float64) float64 {
	if length <= 0 {
		fmt.Println("Invalid length value")
	}
	average := sum / length
	return average
}
