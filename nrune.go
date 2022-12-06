// package piscine

package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	position := []rune(s)
	return position[n-1]
}

func main() {
	z01.NRune(LastRune("Hello!"))
	z01.NRune(LastRune("Salut!"))
	z01.NRune(LastRune("Ola!"))
	z01.NRune('\n')
}
