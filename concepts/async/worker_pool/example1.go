package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker pool definition
type WorkerPool struct {
	tasks       []Task
	concurrency int
	taskChannel chan Task
	wg          sync.WaitGroup
}

// task definition
type Task struct {
	id int
}

// way to process the tasks
func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.id)
	// simulating time consuming operation
	time.Sleep(2 * time.Second)
}

// function to execute the worker pool
func (wp *WorkerPool) worker(index int) {
	for task := range wp.taskChannel {
		fmt.Printf("Worker %d started processing task %d\n", index, task.id)
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// create a buffered channel to hold tasks
	// Will use this channel to pass tasks from main goroutine to worker goroutines
	wp.taskChannel = make(chan Task, len(wp.tasks))

	// start the worker
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker(i)
	}

	// telling wait_group to wait for 20 tasks before reaching the done state
	wp.wg.Add(len(wp.tasks))

	// send tasks to task channel & close the channel after all tasks are sent
	// go routines only start processing tasks when all 20 tasks are sent to the channel
	for _, task := range wp.tasks {
		fmt.Printf("Sending task %d to the pool\n", task.id)
		wp.taskChannel <- task
	}

	fmt.Printf("All tasks have been sent to the pool. Closing the channel now\n")
	close(wp.taskChannel)

	// pause the main goroutine until all tasks are processed
	wp.wg.Wait()
}

func main() {
	// create new tasks slice
	// here we create 20 tasks with ids from 1 to 20
	tasks := make([]Task, 20)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = Task{id: i + 1}
	}

	// create a worker pool
	wp := WorkerPool{
		tasks:       tasks,
		concurrency: 5, // number of workers that can run at a time
	}

	// Run the pool
	wp.Run()
	fmt.Printf("All tasks have been processed")

}
