package main

import (
	"fmt"
	"net/http"
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
	// using the infinite loop
	for {
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	fmt.Println("Starting: " + link)

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}
