package main

// package piscine

import "fmt"

func ConcatParams(args []string) string {
	emptystring := ""
	count := 0

	for range args {
		count++
	}

	for i := range args {
		if i == count-1 {
			emptystring = result + args[i]
		} else {
			emptystring = result + args[i] + "\n"
		}
	}
	return result
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
