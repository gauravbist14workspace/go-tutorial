package main

import "fmt"

type sumCombinationConfig struct {
	arr    []int
	k      int
	result [][]int // to store the output combination
}

func NewSumCombinationConfigConfig(arr []int, k int) *sumCombinationConfig {
	return &sumCombinationConfig{
		arr:    arr,
		k:      k,
		result: make([][]int, 0), // Initialize output_arr to an empty slice
	}
}

func (config *sumCombinationConfig) generateCombinations() {
	config.process(0, 0, []int{})
	fmt.Printf("Combinations that sum to %d: %v\n", config.k, config.result)
}

func (config *sumCombinationConfig) process(index int, sum int, output_arr []int) {
	if sum == config.k {
		fmt.Printf("Found combination: %v\n", output_arr)
		d := make([]int, len(output_arr))
		copy(d, output_arr)                      // Create a copy of the current output array
		config.result = append(config.result, d) // Append the copy to output_arr
		return
	} else if sum > config.k {
		return // If the sum exceeds k, no need to continue
	}

	for i := index; i < len(config.arr); i++ {
		output_arr = append(output_arr, config.arr[i])

		config.process(i, sum+config.arr[i], output_arr)

		output_arr = output_arr[:len(output_arr)-1]
	}
}
