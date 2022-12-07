// package main

package piscine

// import "fmt"

func IsUpper(s string) bool {
	for _, j := range s {
		if !(j >= 'A' && j <= 'Z') {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsUpper("HELLO"))
// 	fmt.Println(IsUpper("HELLO!"))
// }
