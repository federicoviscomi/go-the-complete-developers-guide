package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type englishBot struct {
}
type spanishBot struct {
}
type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (eb englishBot) getGreeting() string {
	return "Hi!"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
func main() {
	sb := spanishBot{}
	printGreeting(sb)
	eb := englishBot{}
	printGreeting(eb)

	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("get error ", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	n, err := response.Body.Read(bs)
	if err == io.EOF {
		fmt.Println("EOF ok")
	} else if err != nil {
		fmt.Println("body read err", err)
	}
	fmt.Println(string(bs))
	fmt.Println("byte read: ", n)
}
