// package main

package piscine

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}

// func add10(x int) {
// 	x = 10
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	ForEach(add10(10), a[3])
// }
