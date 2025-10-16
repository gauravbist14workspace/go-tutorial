package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex // Mutex to protect counter

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			fmt.Println("Updating couter in go-routine 1")
			counter++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			fmt.Println("Updating couter in go-routine 2")
			counter++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value (with mutex):", counter)
}
