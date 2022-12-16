// package main

package piscine

// import "fmt"

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a); i++ {
		if i < len(a) {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) > 0 {
					return false
				}
			}
		} else if i == len(a) {
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		}
	}
	return true
}

// func AscendingOrder(a, b int) int {
// 	if a < b {
// 		return -1
// 	} else if a == b {
// 		return 0
// 	}
// 	return 1
// }

// func main() {
// 	a1 := []int{0, 1, 2, 3, 4, 5}
// 	a2 := []int{0, 2, 1, 3}

// 	result1 := IsSorted(AscendingOrder, a1)
// 	result2 := IsSorted(AscendingOrder, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
