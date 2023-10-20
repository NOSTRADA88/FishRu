package functions

import "unicode"

func ToUpperFirstSymbol(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
