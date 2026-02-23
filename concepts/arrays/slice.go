package main

import (
	"fmt"
)

func slice_practise() {
	// Create a slice from an array
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]  // Slice from index 1 to 3 (exclusive of 4)
	fmt.Println(slice) // Output: [20 30 40]

	// Create a slice using make
	slice2 := make([]int, 3) // Create a slice of length 3
	slice2[0] = 10
	fmt.Println(slice2) // Output: [10 0 0]

	// Append to a slice
	slice2 = append(slice2, 20)
	fmt.Println(slice2) // Output: [10 0 0 20]

	// -------------------------------------------------------------
	// Initiate a 2D slice for let's say dynammic_programming
	rows, cols := 3, 4
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
}
