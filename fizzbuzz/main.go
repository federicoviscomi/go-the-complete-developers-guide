package main

import "fmt"

type wb struct {
	word string
	div  int
}

func main() {
	wordsDiv := []wb{
		{"Fizz", 3},
		{"Buzz", 5},
		{"Srolli", 7},
	}
	for i := 0; i < 35; i++ {
		divisible := false
		for _, wb := range wordsDiv {
			if i%wb.div == 0 {
				fmt.Print(wb.word)
				divisible = true
			}
		}
		if divisible {
			fmt.Println()
		} else {
			fmt.Println(i)
		}
	}
}
