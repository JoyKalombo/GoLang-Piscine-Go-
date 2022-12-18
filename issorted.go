// package main

package piscine

// import "fmt"

func IsSorted(f func(a, b int) int, a []int) bool {
	result := 1
	result1 := 1
	result2 := 1

	for k, v := range a {
		if k != len(a)-1 {
			if f(v, a[k+1]) < 0 {
				result++
			}
			if f(v, a[k+1]) > 0 {
				result1++
			}
			if f(v, a[k+1]) == 0 {
				result2++
			}
		}
	}
	return result == len(a) || result1 == len(a) || result2 == len(a)
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
