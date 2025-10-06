package client

import (
	"fmt"
	"io"
	"net/http"
)

func NewClient() {
	clientConn := &http.Client{}
	fmt.Println("Created a new client object")

	greetReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello/bob", nil)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	resp, err := clientConn.Do(greetReq)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		fmt.Printf("Failed to get a successful response from server. Got: %v", string(bodyBytes))
	}

	limitedReader := io.LimitReader(resp.Body, 1024*1024) // read 1MB max
	bodyBytes, err := io.ReadAll(limitedReader)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("Got response from server:\n%v", string(bodyBytes))
}
