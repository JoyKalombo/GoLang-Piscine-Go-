package piscine

func Swap(a *int, b *int) {

	c := *a
	*a = *b 
	*b = c
	
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := 1
// 	piscine.Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }