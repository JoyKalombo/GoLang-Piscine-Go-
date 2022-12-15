// package main

package piscine

// import "fmt"

func CountIf(f func(string) bool, tab []string) int {
	countinteger := 0
	for _, v := range tab {
		if f(v) {
			countinteger++
		}
	}
	return countinteger
}

// func IsNumeric(s string) bool {
// 	for _, k := range s {
// 		if !(k >= '0' && k <= '9') {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}

// 	answer1 := CountIf(IsNumeric, tab1)
// 	answer2 := CountIf(IsNumeric, tab2)

// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }
