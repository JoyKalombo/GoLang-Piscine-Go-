package piscine

// import (
// 	"fmt"
// 	"piscine"
// )

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

// func main() {
// 	a := 0
// 	b := 1
// 	piscine.Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
