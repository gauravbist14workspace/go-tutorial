package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // Only the sender should close a channel, never the receiver OR else it cause panic.
}

func fibonacciExample1() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // The loop receives values from the channel repeatedly until it's closed.
		fmt.Println(i)
	}
}
