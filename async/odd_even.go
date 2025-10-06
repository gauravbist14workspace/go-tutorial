package main

import (
	"fmt"
	"sync"
)

func OddEvenDemo() {
	chEven := make(chan int)
	chOdd := make(chan int)
	wg := &sync.WaitGroup{}

	// Set the count to 3: one for the sender and one for each receiver
	wg.Add(3)

	go func(ch chan int) {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Even:", val)
		}
	}(chEven)

	go func(ch chan int) {
		defer wg.Done()
		for val := range ch {
			fmt.Println("Odd:", val)
		}
	}(chOdd)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ { // Corrected this line
			if i%2 == 0 {
				chEven <- i
			} else {
				chOdd <- i
			}
		}
		close(chEven)
		close(chOdd)
	}()

	wg.Wait()
	fmt.Println("All nums are printed")
}
