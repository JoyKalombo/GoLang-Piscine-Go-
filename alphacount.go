// package main

package piscine

// import "fmt"

func AlphaCount(s string) int {
	count := 0

	for _, s := range s {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
			count++
		}
	}
	return count
}

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(s)
// 	fmt.Println(nb)
// }
