// package main

package piscine

// import "fmt"

func MakeRange(min, max int) []int {
	if !(min < max) {
		return nil
	} else {
		slice := make([]int, max-min)
		for i := 0; i < (max - min); i++ {
			slice[i] = min + i
		}
		return slice
	}
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
