package main

import (
	"fmt"
	"sync"
)

func factorial(num int, ch chan<- int, wg *sync.WaitGroup) {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	ch <- result
	wg.Done()
}

func main() {
	n := 4

	wg := &sync.WaitGroup{}

	ch := make(chan int)

	for i := n; i > 0; i-- {
		wg.Add(1)
		go factorial(i, ch, wg)
	}

	finalSum := 0
	for i := n; i > 0; i-- {
		finalSum += <-ch
	}

	close(ch)

	wg.Wait()
	fmt.Println(finalSum)
}
