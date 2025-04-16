package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"https://prothomalo.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	// creating a channel
	c := make(chan string)

	for _, link := range links {
		// introducing this go keyword -> run this checklink function in a new go routine for each link
		go checkLink(link, c)
	}

	/*
	// five messages to look at for five different links
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	*/

	/**
	 - The line go checkLink(<-c, c) retrieves a link from the channel c using <-c. 
	 - This means that the program is waiting for a link to be sent to the channel, which could be a link that has already been checked.
	 - After the checkLink function processes the link (whether it finds the link is up or down), it sends a message back through the same channel c.
	 - The infinite loop continuously listens for new links to be sent to the channel. When a link is received, it starts a new goroutine to check that link again


	 for {
	 	go checkLink(<-c, c)
	 }
	*/


	// function literal: anonymous function/lambda function
	// the loop l := range c continuously receives values from channel c. Once c has a value, it is assigned to l
	// The variable link inside anonymous function holds the value of l at the time this function is called. This is important as it captures the current state of l when that goroutine is created
	// if l was directly used, it could lead to unexpected behavior. By assigning l to link, we ensure that we are looking at the current state of l
	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	} 
}

func checkLinkSimple(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up")
	c <- "Yep, it's up"
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
	}
	fmt.Println(link, "is up!")
	c <- link
}