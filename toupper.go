// package main

package piscine

// import (
// 	"fmt"
// )

func IsLower(s string) bool {
	for _, j := range s {
		if !(j >= 'a' && j <= 'z') {
			return false
		}
	}
	return true
}

func ToUpper(s string) string {
	var storagestring string = ""
	for _, i := range s {
		if IsLower(string(i)) {
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
