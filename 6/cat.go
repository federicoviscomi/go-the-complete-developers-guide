package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(1)
	}
	written, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error copying to stdout", err)
		os.Exit(1)
	}
	fmt.Println("\ncopied bytes: ", written)
}
