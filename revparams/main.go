package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for eachargument := range arguments {
		for _, char := range arguments[(len(arguments)-1)-eachargument] {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')
	}
}
