package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	// readFileInOneGo()

	// readFileLineByLine()

	readFileInBytes()

}

func readFileInOneGo() {
	b, err := os.ReadFile("file/dummy.txt")
	if err != nil {
		fmt.Println("Some error while reading file in one go")
		return
	}

	fmt.Println(string(b))
}

func readFileLineByLine() {
	file, err := os.Open("file/dummy.txt")
	if err != nil {
		fmt.Println("Some error while reading file line by line")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		time.Sleep(300 * time.Millisecond)
	}
}

func readFileInBytes() {
	file, err := os.Open("file/dummy.txt")
	if err != nil {
		fmt.Println("Some error while reading file line by line")
		return
	}
	defer file.Close()

	chunk := make([]byte, 32) // read 32 bytes at a time from file

	for {
		n, err := file.Read(chunk)
		if err != nil && err != io.EOF {
			fmt.Println("Some error while reading file line by line")
		}

		if n == 0 {
			fmt.Println("reach the end of file")
			break
		}

		fmt.Printf("%v", string(chunk[:n]))
		time.Sleep(300 * time.Millisecond)
	}
}
