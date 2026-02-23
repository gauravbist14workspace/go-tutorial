package main

import (
	"math/rand"
	"time"
)

func GenerateRandom() string {
	// @deprecated Always seed the random number
	rand.Seed(time.Now().Unix())

	randomString := []string{"START", "PAUSE", "STOP"}[rand.Intn(3)]
	return randomString
}
