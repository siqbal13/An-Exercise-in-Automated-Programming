package main

import "testing"

func TestLinreg(t *testing.T) {
	var x = []float64{1, 2, 3}
	var y = []float64{1, 2, 3}
	var expectedSlope = 1.0
	var expectedIntercept = 0.0
	m, c := linreg(x, y)
	if m != expectedSlope || c != expectedIntercept {
		t.Errorf("linreg(%v, %v) = (%f, %f), want (%f, %f)", x, y, m, c, expectedSlope, expectedIntercept)
	}
}
