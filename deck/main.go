package main

import "fmt"
func main() {
	deck := newDeck()
	const fileName = "./deck.txt"
	err := deck.saveDeckToFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	deckLoadedFromFile, err := loadDeckFromFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	deckLoadedFromFile.print()
	_ = deckLoadedFromFile.saveDeckToFile("./deck_1.txt")
}
