package main

import "fmt"

func createArrayOfSize(size int) []int {
	answer := make([]int, size)

	for i := 0; i < size; i++ {
		answer[i] = i+1
	}
	return answer
}

func main() {
	size := 81
	myarray := createArrayOfSize(size)

	fmt.Println(myarray)
}