package main

import (
	"fmt"
	"os"
)

func main() {
	if !(len(os.Args) <= 1) {
		os.Args = os.Args[1:]
	} else {
		fmt.Printf("File name missing\n")
		return
	}
	if len(os.Args) == 1 {
		file, err := os.Open(os.Args[0])
		if err != nil {
			fmt.Printf("the mistake is : %v\n", err.Error())
		} else {
			arr := make([]byte, 14)
			file.Read(arr)
			fmt.Println(string(arr))
			file.Close()
		}
	} else {
		fmt.Printf("Too many arguments\n")
	}
}
