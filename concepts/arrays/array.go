package main

import (
	"fmt"
)

func array_practise() {
	var arr [5]int   // Declare an array of size 5
	arr[0] = 10      // Assign a value to the first element
	fmt.Println(arr) // Output: [10 0 0 0 0]

	arr2 := [4]string{"Go", "is", "fun"} // Declare and initialize ["Go", "is", "fun", ""]
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}                    // Array with inferred size
	fmt.Printf("%v %v %v", arr3, len(arr3), cap(arr3)) // Output: [1 2 3 4 5]

	// ------------------------------------------------------------------------------------

	// one way to initiate a 2D array
	fixed_2d_array := [3][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

}
