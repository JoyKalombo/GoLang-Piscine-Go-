// package main

package piscine

// import (
// 	"fmt"
// )

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	for _, i := range arg {
		if i >= 0 && i < len(arg)-1 {
			if arg[i] > arg[i+1] {
				arg[i], arg[i+1] = arg[i+1], arg[i]
			}
		} else if i == len(arg)-1 {
			if arg[i-1] > arg[i] {
				arg[i], arg[i-1] = arg[i-1], arg[i]
			}
		}
	}
	return arg[2]
}

// func main() {
// 	middle := Abort(2, 3, 8, 5, 7)
// 	fmt.Println(middle)
// }
