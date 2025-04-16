package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	for _, link := range links {
		// introducing this go keyword -> run this checklink function in a new go routine for each link
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
	}

	fmt.Println(link, "is up")
}