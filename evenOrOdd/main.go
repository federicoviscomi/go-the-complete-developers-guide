package main

import "fmt"

func main() {
	for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
