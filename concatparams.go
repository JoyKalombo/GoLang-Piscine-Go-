// package main

package piscine

// import "fmt"

func ConcatParams(args []string) string {
	emptystring := ""
	count := 0

	for range args {
		count++
	}

	for i := range args {
		if i == count-1 {
			emptystring += args[i]
		} else {
			emptystring += args[i] + "\n"
		}
	}
	return emptystring
}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
