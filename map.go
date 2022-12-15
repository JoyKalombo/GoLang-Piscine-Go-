// package main

package piscine

// import (
// 	"fmt"
// )

// func IsPrime(nbr int) bool {
// 	if nbr <= 1 {
// 		return false
// 	} else if nbr == 2 {
// 		return true
// 	} else {
// 		for i := 2; i < nbr; i++ {
// 			if nbr%i == 0 {
// 				return false
// 			}
// 			return true
// 		}
// 	}
// 	return false
// }

func Map(f func(int) bool, a []int) []bool {
	b := make([]bool, len(a))
	for i := range a {
		b[i] = f(a[i])
	}
	return b // I have to return something here...
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	result := Map(IsPrime, a)
// 	fmt.Println(result)
// }
