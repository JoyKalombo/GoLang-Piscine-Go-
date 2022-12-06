// package main

package piscine

// import (
// 	"fmt"
// )

func Index(s string, toFind string) int {
	str := s
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if j >= i {
				if str[i:j] == toFind {
					return i
				}
			}
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
