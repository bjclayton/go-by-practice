package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://go.dev",
		"https://facebook.com",
		"https://amazon.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for l := range c {
		go func(url string) {
			time.Sleep(time.Second)
			checkUrl(url, c)
		}(l)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
