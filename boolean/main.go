package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayOfStrings := []rune(s)
	for i := 0; i < len(arrayOfStrings); i++ {
		z01.PrintRune(arrayOfStrings[i])
	}
	z01.PrintRune('\n')
}

func even(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func isEven(nbr int) bool {
	if even(nbr) {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) {
		EvenMsg := "I have another even number of arguments"
		printStr(EvenMsg)
	} else {
		OddMsg := "I have an odd number of arguments"
		printStr(OddMsg)
	}
}
