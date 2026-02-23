package main

import (
	"fmt"
)

func main() {
	// Generating random
	trafficSignal := GenerateRandom()
	fmt.Printf("current (random) signal is %v\n", trafficSignal)

	// Comparing 2 numers [MAX MIN]
	a := 104
	b := 105
	max := GetBiggestFrom2(a, b)
	fmt.Printf("biggest from %v and %v is %v\n", a, b, max)
}
