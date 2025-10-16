package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// write code to write into dummy.txt file
	file, err := os.Create("file/dummy2.txt")
	if err != nil {
		fmt.Printf("failed to create file: %s", err)
	}
	defer file.Close()

	file.Write([]byte("hello world"))

	_, err = io.Copy(file, strings.NewReader("\n i hope this is next line"))
	if err != nil {
		fmt.Printf("failed to write into file: %s", err)
	}

}
