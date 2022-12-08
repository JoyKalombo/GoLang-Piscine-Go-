package main

// package piscine

import "fmt"

func TrimAtoi(s string) int {
	num := 0
	count := 0
	list := []rune(s)
	for _, char := range list {
		for i := rune(48); i < char; i++ {
			if i >= '0' && i <= '9' {
				count++
			}
		}
		num = num*10 + count
		count = 0
	}
	return num
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
