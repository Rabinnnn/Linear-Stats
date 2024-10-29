package main

import (
	"bufio"
	"fmt"
	"linear-stats/maths"
	"math"
	"os"
	"strconv"
)

func main() {
	var ysum float64
	var xaxis []float64

	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect number of arguments")
		return
	}

	// capture the data file entered by the user
	inputfilename := os.Args[1]
	inputfile, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer inputfile.Close() // close file when the program finishes

	fileinfo, err := inputfile.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// check if the file has contents
	if fileinfo.Size() == 0 {
		fmt.Println("Error: File is empty")
		return
	}

	scanner := bufio.NewScanner(inputfile)

	var data []float64

	//  read each line from the file, parse it into a float64, and append it to the data slice.
	for scanner.Scan() {
		new := scanner.Text()
		if new == "" {
			continue
		}

		val, err := strconv.ParseFloat(new, 64)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		data = append(data, val)
	}

	yaxis := data
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 0; i < len(data); i++ {
		ysum += data[i]  // calculate sum of all yaxis values
		xaxis = append(xaxis, float64(i)) // use indices as coordinates for each data point
	}

	var xsum float64
	for i := 0; i < len(xaxis); i++ {
		xsum += xaxis[i]
	}

	length := float64(len(data))
	xmean := maths.Average(xsum, length)
	ymean := maths.Average(ysum, length)

	// Calculate Pearson Correlation Coefficient
	numerator := 0.0
	denominatorX := 0.0
	denominatorY := 0.0

	for i := 0; i < len(xaxis); i++ {
		xDiff := xaxis[i] - xmean
		yDiff := yaxis[i] - ymean

		numerator += xDiff * yDiff
		denominatorX += xDiff * xDiff
		denominatorY += yDiff * yDiff
	}

	denominator := math.Sqrt(denominatorX * denominatorY)
	r := numerator / denominator

	// Calculate Linear Regression Coefficients
	a, b := maths.LinearRegression(xaxis, yaxis)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
