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

type logWriter struct {
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
	//readBody(response)

	lw := logWriter{}
	n, e := lw.Write([]byte{0, 0, 0, 0})
	fmt.Println(n)
	fmt.Println(e)

	written, err := io.Copy(os.Stdout, response.Body)
	fmt.Println()
	fmt.Println("byte copied ", written)
	fmt.Println(os.File{})

	s := square{sideLength: 0.5}
	s.printArea()

	t := triangle{height: 0.4, base: 0.6}
	t.printArea()
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func readBody(response *http.Response) {
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
