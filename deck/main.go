package main

import (
	"fmt"
)

func main() {
	d := newDeck()
	hand, remainingDeck := deal(d, 6)
	hand.print()
	remainingDeck.shuffle()
	fmt.Println()
	greeting := "oh hello!"
	fmt.Println(greeting)
	fmt.Println([]byte(d.toString()))
}
