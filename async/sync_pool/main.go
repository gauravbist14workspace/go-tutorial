package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	pool := getSyncPoolInstance(5)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			obj := pool.Get().([]byte) // Get a byte slice from the pool
			fmt.Print("-")

			time.Sleep(100 * time.Millisecond) // Simulate some work

			pool.Put(obj) // Return the byte slice to the pool
			wg.Done()
		}()
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func getSyncPoolInstance(initialCap int) *sync.Pool {

	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Print(".")
			return make([]byte, 1024)
		},
	}

	for i := 0; i < initialCap; i++ {
		pool.Put(make([]byte, 1024))
	}

	return pool
}
