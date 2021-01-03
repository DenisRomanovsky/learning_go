package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://vk.com",
		"http://golan.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // go == run the function in a routine
		// Need channels to show main routine that children are done.
		// Otherwise main will halt before the children are done.
	}

	// As soon as the check is done - send the
	// message to channel -> trigger the check again
	for l := range c {
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	fmt.Println("Starting: " + link)

	_, err := http.Get(link)
	if err != nil {
		logWithRestart(" might be down!", link, c)
		return
	}

	logWithRestart(" is up!", link, c)
}

func logWithRestart(message string, link string, c chan string) {
	fmt.Println(message + link)
	time.Sleep(time.Second * 2) // Sleep the routine for 2 seconds
	c <- link
}
