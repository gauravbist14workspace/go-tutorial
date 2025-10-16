package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	chOdd := make(chan int)
	chEven := make(chan int)

	wg.Add(3)
	go func(ch chan int) {
		for val := range ch {
			fmt.Println("Odd:", val)
		}
		wg.Done()
	}(chOdd)

	go func(ch chan int) {
		for val := range ch {
			fmt.Println("Even:", val)
		}
		wg.Done()
	}(chEven)

	go func() {
		for i := range 10 {
			if i%2 == 0 {
				chEven <- i
			} else {
				chOdd <- i
			}
		}

		close(chOdd)
		close(chEven)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All done!")

}
