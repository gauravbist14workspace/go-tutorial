package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func(ch chan<- int) {
		time.Sleep(2 * time.Second)
		ch <- 2

		// time.Sleep(10 * time.Second)
		// ch <- 10
	}(ch)

	go func(ch chan<- int) {
		time.Sleep(4 * time.Second)
		ch <- 4

		// time.Sleep(10 * time.Second)
		// ch <- 10
	}(ch)

	timeout := time.After(6 * time.Second)

	// for { // not adding N = 2 here would lead the loop to run till the end of timeout even though the API responses is already present
	// 	select {
	// 	case <-timeout:
	// 		fmt.Println("timeout has occured")
	// 		return
	// 	case res1 := <-ch:
	// 		fmt.Printf("received %d from channel\n", res1)
	// 	case res1 := <-ch:
	// 		fmt.Printf("received %d from channe2\n", res1)
	// 	}
	// }

	// this below one is INCORRECT way to handle multi-api call with timeout
	// If 1 or even both APIs are timed out, it stil enters the default loop & keep on waiting indefinitely for them
	for {
		select {
		case <-timeout:
			fmt.Println("timeout has occured")
			return
		default:
			res1 := <-ch
			res2 := <-ch

			fmt.Printf("received %d from channel\n", res1)
			fmt.Printf("received %d from channe2\n", res2)
		}
	}
}
