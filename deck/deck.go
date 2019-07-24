package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
	"time"
)

type deck []string

const separator = ","

func equals(left deck, right deck) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	for index, leftCard := range left {
		if leftCard != right[index] {
			fmt.Println("cards at index ", index, "differ")
			fmt.Println("\""+leftCard+"\"", "\""+right[index]+"\"")
			return false
		}
	}
	return true
}

func newDeck() deck {
	deck := deck{}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suites := []string{"♣ clubs", "♦ diamonds", "♥ hearts", "♠ spades"}
	for _, suite := range suites {
		for _, value := range values {
			deck = append(deck, strings.TrimSpace(value)+" of "+strings.TrimSpace(suite))
		}
	}
	return deck
}

func (d deck) toString() string {
	return strings.Join(d, separator)
}

func (d deck) shuffleLinear() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		j := rand.Intn(len(d)-i) + i
		fmt.Println("swapping ", i, " and ", j)
		d[i], d[j] = d[j], d[i]
	}
}

func randomInteger(startRangeIncluded int, endRangeExcluded int) int {
	length := endRangeExcluded - startRangeIncluded;
	return int(math.Floor(rand.Float64()*float64(length))) + startRangeIncluded
}

func (d deck) shuffle_nSquare() {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range d {
		for j, _ := range d {
			if randomDecision() {
				d[i], d[j] = d[j], d[i]
			}
		}
	}
}

func randomDecision() bool {
	return rand.Float32() < 0.5
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, ":", card)
	}
}

func (d deck) toByteArray() []byte {
	return []byte(d.toString())
}

func (d deck) saveDeckToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteArray(), 0777)
}

func (d deck) length() int {
	return len(d)
}

func loadDeckFromFile(filename string) (deck, error) {
	deckByteArray, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return deck(strings.Split(string(deckByteArray), separator)), nil
}
