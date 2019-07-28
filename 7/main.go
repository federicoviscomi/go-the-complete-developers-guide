package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	fmt.Println(c)
	for _, link := range links {
		go checkLink(link, c)
	}
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}
	for l := range c {
		//go checklink2(l, c)
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checklink2(link string, c chan string) {
	time.Sleep(time.Second)
	checkLink(link, c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if nil == err {
		fmt.Println(link, " ok")
		c <- link
	} else {
		fmt.Println(link, " err")
		c <- link
	}
}
