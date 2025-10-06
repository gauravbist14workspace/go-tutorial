package channel

import (
	"fmt"
	"time"
)

func BufferedExample() {

	ch := make(chan int, 2) // we can't push new value to channel when it's FULL

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 500
	}()
	fmt.Printf("First value (500): %v\n", <-ch)

	ch <- 100

	time.Sleep(3 * time.Second) // this will hold the printing line even though 100 is pushed to channel until 3 seconds pass
	ch <- 200
	// ch <- 300 // uncommenting this LINE would lead to error since our channel has max buffer set to 2 right now
	fmt.Printf("Next values: %v and %v", <-ch, <-ch)
}
