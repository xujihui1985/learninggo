package receiver

import (
	"strings"
)

type F func(string) string

func (f F) upperCase(s string) string {
	return strings.ToUpper(f(s))
}

func (f F) lowerCase(s string) string {
	return strings.ToLower(f(s))
}

func doubleString(s string) string {
	return s + s
}
func initial(s string) string {
	return string([]rune(s)[0])
}
