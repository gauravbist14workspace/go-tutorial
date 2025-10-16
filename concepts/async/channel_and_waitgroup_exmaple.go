package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WorkInChannels(timeInSeconds time.Duration, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure Done is called even if the function panics
	fmt.Println("Starting a new job")
	time.Sleep(timeInSeconds)
	fmt.Println("Finished the job")

	ch <- fmt.Sprintf("Result of job is %v", rand.Intn(100))
}

func channelAndWaitGroupExample() {
	start := time.Now()

	resultCh := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go WorkInChannels(2*time.Second, resultCh, wg)
	go WorkInChannels(4*time.Second, resultCh, wg)

	// NOTE -> We can use the defer function to execute at the end of the main function to print the total time taken for all jobs to complete.
	defer func() { fmt.Printf("All jobs completed in %v seconds\n", time.Since(start)) }()

	go func() {
		for res := range resultCh {
			fmt.Println(res)
		}
		// fmt.Printf("All jobs completed in %v seconds\n", time.Since(start)) // NOTE -> To print it, we'd have to use sleep at end, to hold the main goroutine, otherwise it exits before this line is executed
	}()

	wg.Wait()
	close(resultCh) // this will STOP the above FOR LOOP & actually finish the 'annonymous' go routine
	// time.Sleep(time.Second) // NOTE -> this HACK will let me print final print statement by holding the 'main' go routine to exit

}
