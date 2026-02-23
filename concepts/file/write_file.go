package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteIntoFile(filePath string) {
	// write code to write into dummy.txt file
	file, err := os.Create(filePath)
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
