package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args

	for _, i := range arguments[0] {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			z01.PrintRune(i)
		}
	}
}
