// package main

package piscine

// import "fmt"

func IsNumeric(s string) bool {
	for _, j := range s {
		if !(j >= '0' && j <= '9') {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsNumeric("010203"))
// 	fmt.Println(IsNumeric("01,02,03"))
// }
