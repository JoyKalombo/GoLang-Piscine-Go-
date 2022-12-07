// package main

package piscine

// import "fmt"

func IsAlpha(s string) bool {
	for _, j := range s {
		if !((j >= 'a' && j <= 'z') || (j >= 'A' && j <= 'Z') || (j >= '0' && j <= '9')) {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("Hello! How are you?"))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha("Whatsthis4"))
// }
