// package main

package piscine

// import "fmt"

func AppendRange(min, max int) []int {
	var array []int = nil

	if max <= min {
		array = nil
	} else {
		array = []int{}

		for i := min; i < max; i++ {
			array = append(array, i)
		}
		return array
	}
	return nil
}

// func main() {
// 	fmt.Println(AppendRange(5, 10))
// 	fmt.Println(AppendRange(10, 5))
// }
