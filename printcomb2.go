// package main

package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			if !(i == '9' && j == '9') {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else if i == '9' && j == '9' {
				z01.PrintRune(i)
				z01.PrintRune(j)
			}
		}
	}
}

// func main() {
// 	PrintComb2()
// }
