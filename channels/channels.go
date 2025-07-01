package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {

	// channels()
	// multiChannels()
	channelsWithClose()
}

// sending single data in single channel in go routines
func channels() {

	ch := make(chan string)

	go func() {
		ch <- "abc"
	}()

	fmt.Println(<-ch)

}

// sending more data in single channels and receiving
func multiChannels() {
	ch := make(chan string)
	a := [5]int{1, 2, 3, 4, 5}

	for _, v := range a {

		v1 := cast.ToString(v)
		go func() {
			ch <- v1
		}()

	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // Here we are waiting for extra data in this channel which is never passed,
	// this will throw error stating that all routines are sleep will not execute
	// user close(ch) function to avoid this error
}

// sending and receiving multiple data in go routines with close channel functionality
// if channels are closed, we cannot send data but can still receive
// This will avoid unecessary waiting of data received from channels
func channelsWithClose() {
	ch := make(chan string)

	a := [5]int{1, 2, 3, 4, 5}
	go func() {
		for _, v := range a {
			v1 := cast.ToString(v)
			ch <- v1
		}
		close(ch) // channels should be closed only when all data are sent, do not close in between result in error "send on close channel"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // this will not throw error even if we are waiting for unsent data through channels
	fmt.Println(<-ch)
}
