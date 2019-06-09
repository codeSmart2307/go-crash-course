package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}

	fmt.Println(<- c)
}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("There was a problem accessing", link, "with error: ", err)
		c <- "Server might be down"
	}

	fmt.Println("The server at", link, "responded with 200 OK")
	c <- "Server is up and running"
}
