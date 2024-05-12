// Here I'm using Chat GPT to generate GO code. My prompt was 'Generate Go code for linear regression analysis on the Anscombe data set. Make it clean, optimized and readable. '.

package main

import (
	"fmt"
	"math"
)

// Define a struct for holding x and y values
type Point struct {
	X float64
	Y float64
}

// Function to calculate mean
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Function to calculate variance
func variance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-mean, 2)
	}
	return sum / float64(len(data))
}

// Function to calculate covariance
func covariance(x []float64, meanX float64, y []float64, meanY float64) float64 {
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += (x[i] - meanX) * (y[i] - meanY)
	}
	return sum / float64(len(x))
}

// Function to perform linear regression
func linearRegression(x []float64, y []float64) (float64, float64) {
	meanX := mean(x)
	meanY := mean(y)

	// Calculate coefficients
	slope := covariance(x, meanX, y, meanY) / variance(x, meanX)
	intercept := meanY - slope*meanX

	return slope, intercept
}

func main() {
	// Anscombe dataset
	anscombe := []Point{
		{10.0, 8.04},
		{8.0, 6.95},
		{13.0, 7.58},
		{9.0, 8.81},
		{11.0, 8.33},
		{14.0, 9.96},
		{6.0, 7.24},
		{4.0, 4.26},
		{12.0, 10.84},
		{7.0, 4.82},
		{5.0, 5.68},
	}

	// Extract x and y values
	var xs []float64
	var ys []float64
	for _, point := range anscombe {
		xs = append(xs, point.X)
		ys = append(ys, point.Y)
	}

	// Perform linear regression
	slope, intercept := linearRegression(xs, ys)

	// Print the results
	fmt.Printf("Slope: %.2f\n", slope)
	fmt.Printf("Intercept: %.2f\n", intercept)
}
