package tools

import "strings"

// Concat take any string in parameter to string.build a new string result of the concatenation
func Concat(keys ...string) string {
	var b strings.Builder
	for i := 0; i < len(keys); i++ {
		b.WriteString(keys[i])
	}
	return b.String()
}
