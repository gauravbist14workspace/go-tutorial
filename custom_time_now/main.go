package main

import (
	"fmt"
	"time"
)

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now()
}

func GenerateClock(clock Clock) string {
	return clock.Now().String()
}

func main() {
	fmt.Println(GenerateClock(RealClock{}))
}
