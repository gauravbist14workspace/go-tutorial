package main

import (
	"fmt"
	"time"
)

func Greet(greeting string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(greeting)
	}
}

func goRoutineDemo() {
	go Greet("World")
	Greet("Hello")
}
