// package main

package piscine

// import (
// 	"fmt"
// )

func ToLower(s string) string {
	var storagestring string = ""
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			storagestring = storagestring + string(rune(i+32))
		} else {
			storagestring = storagestring + string(rune(i))
		}
	}
	return storagestring
}

// func main() {
// 	fmt.Println(ToLower("Hello! How are you?"))
// }
