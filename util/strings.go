package util

import "unicode"

func ToLowerFirstLetter(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func ToUpperFirstLetter(str string) string {
	a := []rune(str)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
