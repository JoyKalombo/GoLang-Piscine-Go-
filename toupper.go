package main

// package piscine

import "github.com/01-edu/z01"

func ToUpper(s string) string {
	for _, s := range s {
		if s == 'a' {
			z01.PrintRune('A')
		} else if s == 'b' {
			z01.PrintRune('B')
		} else if s == 'c' {
			z01.PrintRune('C')
		} else if s == 'd' {
			z01.PrintRune('D')
		} else if s == 'e' {
			z01.PrintRune('E')
		} else if s == 'f' {
			z01.PrintRune('F')
		} else if s == 'g' {
			z01.PrintRune('G')
		} else if s == 'h' {
			z01.PrintRune('H')
		} else if s == 'i' {
			z01.PrintRune('I')
		} else if s == 'j' {
			z01.PrintRune('J')
		} else if s == 'k' {
			z01.PrintRune('K')
		} else if s == 'l' {
			z01.PrintRune('L')
		} else if s == 'm' {
			z01.PrintRune('M')
		} else if s == 'n' {
			z01.PrintRune('N')
		} else if s == 'o' {
			z01.PrintRune('O')
		} else if s == 'p' {
			z01.PrintRune('P')
		} else if s == 'q' {
			z01.PrintRune('Q')
		} else if s == 'r' {
			z01.PrintRune('R')
		} else if s == 's' {
			z01.PrintRune('S')
		} else if s == 't' {
			z01.PrintRune('T')
		} else if s == 'u' {
			z01.PrintRune('U')
		} else if s == 'v' {
			z01.PrintRune('V')
		} else if s == 'w' {
			z01.PrintRune('W')
		} else if s == 'x' {
			z01.PrintRune('X')
		} else if s == 'y' {
			z01.PrintRune('Y')
		} else if s == 'z' {
			z01.PrintRune('Z')
		} else {
			z01.PrintRune(s)
		}
	}
	return ""
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
