package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func doWork(delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting work")
	time.Sleep(delay)
	fmt.Println("Work completed")
}

func WaitGroupDemo() {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	start := time.Now()
	go doWork(2*time.Second, wg)
	go doWork(4*time.Second, wg)

	wg.Wait() // program execution is waiting at this point for 2 tasks to finish
	fmt.Printf("Jobs completed in %v seconds", time.Since(start))

}
