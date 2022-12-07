// package main

package piscine

// import (
// 	"fmt"
// )

func ToUpper(s string) string {
	var storagestring string = ""
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			storagestring = storagestring + string(rune(i-32))
		} else {
			storagestring = storagestring + string(rune(i))
		}
	
	}
	return storagestring
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
