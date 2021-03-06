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

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "responded with error ", err)
		return
	}
	fmt.Println(link, "statusCode: ", resp.StatusCode)
}
