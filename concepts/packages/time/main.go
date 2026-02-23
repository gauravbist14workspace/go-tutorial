package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program starts in 2 seconds...")
	time.Sleep(2 * time.Second)

	ticker := time.NewTicker(5 * time.Second)
	timer := time.NewTimer(15 * time.Second)

	fmt.Printf("starting our timer program at: %v\n", time.Now())
	for {
		select {
		case <-ticker.C:
			fmt.Printf("[%v]\n", time.Now())
		case <-timer.C:
			fmt.Printf("Finished at %v\n", time.Now())
			return
		}
	}
}
