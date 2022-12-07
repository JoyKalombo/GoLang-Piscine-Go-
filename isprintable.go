// package main

package piscine

// import "fmt"

func IsPrintable(s string) bool {
	for _, j := range s {
		if !(int(j) >= 20 && int(j) <= 127) {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPrintable("Hello"))
// 	fmt.Println(IsPrintable("Hello\n"))
// }
