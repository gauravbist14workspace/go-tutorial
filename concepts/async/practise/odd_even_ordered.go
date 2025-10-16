package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int) {
		for val := range ch {
			fmt.Println("Odd: ", val)
			ch <- val + 1
		}
		wg.Done()
	}(ch)

	go func() {
		for i := 1; i <= 10; i += 2 {
			ch <- i
			fmt.Println("Even: ", <-ch)
		}
		close(ch)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all work done")
}
