package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type AlertType string

const (
	BatteryOut    AlertType = "BATTERY_OUT"
	ConsumableOut AlertType = "CONSUMABLE_OUT"
)

type Alert struct {
	AlertId    string
	DeviceName string
	Type       AlertType
	Date       time.Time
}

func main() {

	// readJSONFileInOneGo()

	// readFileLineByLine()

	readFileInBytes()

}

func readJSONFileInOneGo() {
	b, err := os.ReadFile("file/alert.json")
	if err != nil {
		fmt.Println("Some error while reading file in one go")
		return
	}

	// fmt.Println(string(b))

	var alert Alert
	err = json.Unmarshal(b, &alert)
	if err != nil {
		fmt.Println("failed to unmarshall the json with err: ", err)
	}

	fmt.Printf("read the following json from file: %v", alert)
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
