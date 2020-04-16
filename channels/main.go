package main

import (
	"fmt"
	"net/http"
)

func main()  {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<- c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- "something went wrong"
		return
	}

	c <- "all good"
}