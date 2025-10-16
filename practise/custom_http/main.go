package main

import (
	"io"
	"os"
)

func readLineByLine(file io.ReadCloser) {
	str := ""
	chunk := make([]byte, 32)
	for {
		n, err := file.Read(chunk)
		if err != nil && err != io.EOF {

		}

		if err == io.EOF || n == 0 {
			break
		}

		for _, item := range chunk {

		}

	}
}

func main() {
	file, err := os.Open("../../shared/assets/dummy.txt")
	if err != nil {

	}
	defer file.Close()

	readLineByLine(file)

}
