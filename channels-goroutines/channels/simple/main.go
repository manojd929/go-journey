package main

import (
	"fmt"
	"net/http"
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

	// Haha imp concept or loop
	// because go just waits for a value in the channel and the program exits after that if there were no for loop
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		c <- link + " might be down"
		return
	}
	c <- link + " Status Code " + string(resp.Status)
}
