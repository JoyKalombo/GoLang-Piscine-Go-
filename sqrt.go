package piscine

// package main

// import "fmt"

func Sqrt(nb int) int {
	result := nb
	for i := 0; i <= nb; i++ {
		if i*i == result {
			return i
		} else if i*i > nb {
			return 0
		}
	}
	return 0
}

// func main() {
// 	fmt.Println(Sqrt(4))
// 	fmt.Println(Sqrt(3))
// }
