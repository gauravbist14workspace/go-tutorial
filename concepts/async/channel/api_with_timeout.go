package channel

import (
	"fmt"
	"time"

	model "go_tutorial/concepts/async/channel/models"
)

func APICall1(ch chan<- model.APIResponse) {
	res := model.APIResponse{
		Data: nil,
		Err:  fmt.Errorf("there was an error from API 1"),
	}

	fmt.Println("Returning error from API1...")
	time.Sleep(4 * time.Second)
	fmt.Println("Passing Api1 data to channel !")
	ch <- res
}

func APICall2(ch chan<- model.APIResponse) {
	res := model.APIResponse{
		Data: 2,
		Err:  nil,
	}

	fmt.Println("Returning data from API2...")
	time.Sleep(20 * time.Second)

	fmt.Println("Passing Api2 data to channel !")
	ch <- res
}

func TimeOutDemo() {

	// created 2 separate channels to segregate the API responses
	ch1 := make(chan model.APIResponse, 1)
	ch2 := make(chan model.APIResponse, 1)

	// Called the asynchronous APIs (eg: making serice layer calls from experience layer)
	go APICall1(ch1)
	go APICall2(ch2)

	// Need to define timeout outside the FOR LOOP or it keeps on getting reset
	timeout := time.After(time.Second * model.CONNECTION_TIMEOUT)

	// THIS IS WRONG APPROACH:
	// bcoz if timeout condition is not met immediately (which is unlikely to happen), it would automatically go to DEFAULT case
	// WHAT HAPPEN NOW ? Now instead of waiting for timeout in select, we're blocking outside the select inside the DEFAULT case
	/*for {
		select {
		case <-timeout:
			fmt.Print("Connection timed out while waiting for response")
			return
		default:
			x := <-ch1
			y := <-ch2

			if x.Err != nil {
				fmt.Println("failed to get response from 1st api")
			} else {
				fmt.Println("Response from 1st API", x.Data)
			}

			if y.Err != nil {
				fmt.Println("failed to get response from 2nd api")
			} else {
				fmt.Println("Response from 2nd API", y.Data)
			}
		}
	}*/

	// FOR LOOP runs for 2 times only & then moves the control ahead.
	// If we run infinite loop, we won't reach the last line of code w/o triggering the timeout
	// It also breaks if timeout expires
	for range 2 {
		select {
		case <-timeout:
			fmt.Println("Connection timed out while waiting for response")
			return
		case res := <-ch1:
			if res.Err != nil {
				fmt.Printf("API1 failed with error: %v\n", res.Err.Error())
			} else {
				fmt.Printf("Response from API 1: %v\n", res.Data.(int))
			}
		case res := <-ch2:
			if res.Err != nil {
				fmt.Printf("API2 failed with error: %v\n", res.Err.Error())
			} else {
				fmt.Printf("Response from API 2: %v\n", res.Data)
			}
		}
	}

	fmt.Println("Do something after we receive the responses from both API calls")

}
