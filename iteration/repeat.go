package iteration

import "strings"

// Repeat returns repeated character
func Repeat(c string, repeatCount int) (repeat string) {
	return strings.Repeat(c, repeatCount)
}
