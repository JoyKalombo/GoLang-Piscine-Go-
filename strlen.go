package piscine

import "unicode/utf8"

func StrLen(s string) int {
	a := utf8.RuneCountInString(s)
	return a
}
