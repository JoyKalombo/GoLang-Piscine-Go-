package piscine

// package main

// import (
// 	"fmt"
// )

func IterativeFactorial(nb int) int {
	result := 1

	for i := 1; i <= nb; i++ {
		if nb == 0 {
			return 1
		} else if nb < 0 {
			return 0
		} else {
			result = result * i
		}
	}
	return result
}

// func main() {
// 	arg := 4
// 	fmt.Println(IterativeFactorial(arg))
// }
