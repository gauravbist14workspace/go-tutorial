package main

import (
	"fmt"
	"math/bits"
)

/*
* Inputs: 9: 1001 and 10: 1010
* XOR operation: 1001 ^ 1010 = 0011 , XOR returns 1 only when different bits
* Ouput: 2 bits are different
 */
func main() {
	finalByte := uint(9 ^ 10)
	fmt.Printf("1001 ^ 1010 = %04b\n", finalByte) // formatting: padding binary output with 0s until we have 4 characters
	fmt.Printf("bits.OneCount = %v\n", bits.OnesCount(finalByte))
	fmt.Printf("custom solution: %v\n", customBitChecker(9^10))
}

func customBitChecker(value int) int {
	count := 0
	for value > 0 {
		count += value & 1
		value >>= 1
	}

	return count
}
