package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for _, word := range arguments {
		for _, char := range word {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')
	}
}
