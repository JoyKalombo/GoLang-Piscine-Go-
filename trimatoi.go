package piscine

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
