package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkLink(link string, c chan string) {
	// sleep is also not recomended, as api role may be is to check instantly for others use case
	// so use anonymous function
	// time.Sleep(time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Link %v is not working\n", link)
		//c <- "Link is down"
		c <- link
		return
	}
	fmt.Println(link, "is up")
	//c <- "Link is up"

	c <- link

}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://programiz.com",
		"https://geeksforgeeks.org",
	}

	c := make(chan string)
	fmt.Println("Channel is", c)

	for _, link := range links {
		go checkLink(link, c)
	}
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//fmt.Println(<-c)
	//time.Sleep(2 * time.Second)

	// for {
	// 	go checkLink(<-c, c)
	// }

	// ranging over channel, watchover channel whenever value comes assign to l and attempt new go routines
	for l := range c {
		// time.Sleep(time.Second) // main routines wait for each second every time
		fmt.Println("Total goroutines is ", runtime.NumGoroutine())
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		//go checkLink(l, c)
	}

}
