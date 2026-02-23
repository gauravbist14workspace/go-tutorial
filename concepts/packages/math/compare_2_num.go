package main

import "math"

func GetBiggestFrom2(a, b int) int {
	// most math package functions run of float64
	res := math.Max(float64(a), float64(b))
	return int(res)
}
