package functions

import "unicode"

func ToUpperFirstSymbol(str string) string {
	if len(str) <= 0 {
		return ""
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
