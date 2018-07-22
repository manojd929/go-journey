package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// stylistic concern addressed
	// for better readability
	for l := range c {
		// go checkLink(l, c)

		// function literal
		go func(link string) {
			// time.Sleep(d Duration) pauses the current goRoutine
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// we did not add time.Sleep here because we do not want initial pause i.e., for first round of execution
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
