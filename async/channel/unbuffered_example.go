package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func WorkInChannels(delay time.Duration, ch chan<- int) {
	fmt.Printf("Starting a new job\n")
	time.Sleep(delay)

	ch <- rand.Intn(100)
}

func UnBufferedExample() {
	start := time.Now()

	resultCh := make(chan int) // this is an Unbuffered channel

	go WorkInChannels(2*time.Second, resultCh)
	go WorkInChannels(4*time.Second, resultCh)

	res1 := <-resultCh
	fmt.Printf("%v\n", res1)
	res2 := <-resultCh
	fmt.Printf("%v\n", res2)

	fmt.Printf("All jobs completed in %v seconds", time.Since(start))
}
