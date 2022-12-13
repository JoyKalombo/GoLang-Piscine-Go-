package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func numRune(num int) {
	r := '0'
	if num == 0 {
		z01.PrintRune(r)
	}
	for i := 1; i <= num%10; i++ {
		r++
	}
	if num/10 != 0 {
		numRune(num / 10)
	}
	z01.PrintRune(r)
}

func main() {
	points := point{}

	setPoint(&points)
	x := points.x
	y := points.y
	message1 := "x = "
	message2 := ", y = "

	for _, i := range message1 {
		z01.PrintRune(i)
	}

	numRune(x)
	for _, i := range message2 {
		z01.PrintRune(i)
	}
	numRune(y)
	z01.PrintRune('\n')

	// fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
