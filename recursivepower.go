package piscine

// package main

// import (
// 	"fmt"
// )

func RecursivePower(nb int, power int) int {
	if nb == 0 || power < 0 {
		return 0
	} else if nb == 1 || power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

// func main() {
// 	fmt.Println(RecursivePower(4, 3))
// }
