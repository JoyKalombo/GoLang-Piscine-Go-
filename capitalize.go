package main

// package piscine

import "fmt"

func Capitalize(s string) string {
	var storagestring string = ""
	count := 0

	for _, i := range s {
		if i == ' ' {
			storagestring = storagestring + string(rune(i))
			count = 0
		} else if count == 0 && (i >= 'A' && i <= 'Z') {
			storagestring = storagestring + string(rune(i))
			count++
		} else if count == 0 {
			storagestring = storagestring + string(rune(i-32))
			count++
		} else if count != 0 && (i >= 'A' && i <= 'Z') {
			storagestring = storagestring + string(rune(i+32))
		} else {
			storagestring = storagestring + string(rune(i))
			count++
		}
	}

	return storagestring
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
