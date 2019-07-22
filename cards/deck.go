package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

const separator = ","

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, ":", card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"clubs ♣", "diamonds ♦", "hearts ♥", "spades ♠"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join(d, separator)
}

func (d deck) toByteArray() []byte {
	return []byte(d.toString())
}

func (d deck) saveDeckToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteArray(), 0777)
}

func loadDeckFromFile(filename string) (deck, error) {
	deckByteArray, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return deck(strings.Split(string(deckByteArray), separator)), nil
}
