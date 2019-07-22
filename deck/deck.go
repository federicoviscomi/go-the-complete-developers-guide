package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	deck := deck{}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suites := []string{"♣ clubs", "♦ diamonds", "♥ hearts", "♠ spades"}
	for _, suite := range suites {
		for _, value := range values {
			deck = append(deck, fmt.Sprintf("%5.5s of %s", value, suite))
		}
	}
	return deck
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) print() {
	for _, card := range d {
		fmt.Print(card)
	}
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range d {
		for j, _ := range d {
			if (rand1()) {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
}

func rand1() bool {
	return rand.Float32() < 0.5
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
