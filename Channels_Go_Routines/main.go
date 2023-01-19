package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{"http://google.com", "http://facebook.com", "http://golang.org"}
	c := make(chan string)
	for _, site := range websites {
		go checkLink(site, c)
	}
	// for range websites {
	// 	fmt.Println(<-c)
	// }

	// continuous loop through channel receiving messages
	for l := range c {
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(site string, c chan string) {
	resp, err := http.Get(site) // Blocking Call Here
	if err != nil {
		fmt.Println(resp.Status, "Unable to connect to", site)
		c <- site
		return
	}
	fmt.Println(resp.Status, "Received response from", site)
	c <- site
}

// currently the program is being executed exactly in order of the slice of strings.
// Could be an issue if we have a long response time from a single site

// When we execute our program it is actually a Go Routine itself

// Go Scheduler attempts to use only one CPU by default and decides Go Routines to be run
// Concurrency != Parellelism
// Concurrency - We can have multiple threads executing code. If one thread blocks, another is picked up

// Parallelism - Multiple threads executed at the exact same time. Requires multiple cpu/cores

// Channels are used to communicate between Go Routines
// Data that is passed into a Channel is typed!
// c := make(chan string)

// sending data with channels
// channel <- 5 // Send the value 5 into this channel
// myNumber <- channel // wait for a value to be sent into this channel. when we get one, assign the value to 'myNumber'
// fmt.Println(<- channel) Wait for a value to be sent to this channel, when we get one log it out immediately
// Waiting to receive a message from a channel is a blocking code
