// package piscine

package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {
	result := nb
	if nb == 1 || (nb != 0 && power == 0) {
		return 1
	} else if (nb == 0 && power != 0) || power < 0 {
		return 0
	} else if nb == 0 && power == 0 {
		fmt.Println("error")
	} else {
		for i := 1; i <= power-1; i++ {
			result = result * nb
		}
	}
	return result
}

func main() {
	fmt.Println(IterativePower(0, 0))
}
