package utils

import (
	"strings"
	"unicode"
)

func Capitalize(word string) string {
	if len(word) == 0 {
		return ""
	}
	return string(unicode.ToUpper(rune(word[0]))) + word[1:]
}

func ToCamelCase(s string) string {
	parts := strings.Fields(s)
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func ToSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "_"))
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
