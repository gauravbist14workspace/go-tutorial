package main

import "fmt"

func main() {
	// testing_swtich(1)
	// testing_swtich(2)
	// testing_swtich(4)
	// testing_swtich(5)
	testing_swtich(3)
}

func testing_swtich(num int) {
	switch num {
	case 1, 2:
		fmt.Println("received value", num)
	case 3:
	case 4:
		fmt.Println("received value", num)
	default:
		fmt.Print("unknown value receieved")
	}
}
