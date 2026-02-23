package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Hello,    world from GO!"
	reg := regexp.MustCompile(`\s+`)
	confirm := reg.Split(str, -1) // -1 indicates return the entire array. Positive value returns only that amount of elements
	fmt.Printf("%v", confirm)
}
