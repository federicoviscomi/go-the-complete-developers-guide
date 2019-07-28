package main

import "fmt"

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
}
