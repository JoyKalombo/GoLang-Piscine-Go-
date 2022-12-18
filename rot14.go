// package main

package piscine

import "github.com/01-edu/z01"

func Rot14(s string) string {
	for _, i := range s {
		if (i >= rune(65) && i <= rune(76)) || (i >= rune(97) && i <= rune(108)) {
			z01.PrintRune(i + 14)
		} else if (i >= rune(77) && i <= rune(90)) || (i >= rune(109) && i <= rune(122)) {
			z01.PrintRune(i + 14 - 26)
		} else {
			z01.PrintRune(i)
		}
	}
	return ""
}

// func main() {
// 	result := Rot14("Hello! How are You?")

// 	for _, r := range result {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }
