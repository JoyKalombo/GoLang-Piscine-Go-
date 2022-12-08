package piscine

func BasicAtoi2(s string) int {
	num := 0
	count := 0
	list := []rune(s)
	for _, char := range list {
		if char >= '0' && char <= '9' {
			for i := '0'; i < char; i++ {
				count++
			}

			num = num*10 + count
			count = 0

		} else {
			return 0
		}
	}
	return num
}
