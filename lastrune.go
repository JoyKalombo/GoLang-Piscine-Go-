package piscine

// package main

// import "github.com/01-edu/z01"

func LastRune(s string) rune {
	position := []rune(s)
	return position[len(s)-1]
}

// func main() {
// 	z01.PrintRune(LastRune("Hello!"))
// 	z01.PrintRune(LastRune("Salut!"))
// 	z01.PrintRune(LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }
