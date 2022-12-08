// package main

package piscine

// import "fmt"

func BasicAtoi(s string) int {
	num := 0
	count := 0
	list := []rune(s)
	for _, char := range list {
		for i := '0'; i < char; i++ {
			count++
		}
		num = num*10 + count
		count = 0
	}
	return num
}

// func main() {
// 	fmt.Println(BasicAtoi("12345"))
// 	fmt.Println(BasicAtoi("0000000012345"))
// 	fmt.Println(BasicAtoi("000000"))
// }
