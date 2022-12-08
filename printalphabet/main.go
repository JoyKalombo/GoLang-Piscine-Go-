package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
// func main() {
// 	fmt.Println(int('8' + '0'))
// 	fmt.Println(int((rune(8))))
// 	z01.PrintRune(rune(56))
// 	z01.PrintRune('\n')
// }
