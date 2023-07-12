package str

import (
	"fmt"
	"strings"
)

// SurroundAll surrounds all strings in a list with two other strings
func SurroundAll(list []string, left, right string) (result []string) {
	result = make([]string, len(list))
	for i, str := range list {
		result[i] = Surround(str, left, right)
	}
	return result
}

// Surround surrounds a string with two other strings
func Surround(str, left, right string) string {
	return fmt.Sprintf("%s%s%s", left, str, right)
}

// TrimAll trims all strings in a list with a cutset
func TrimAll(list []string, cutset string) (result []string) {
	result = make([]string, len(list))
	for i, str := range list {
		result[i] = strings.Trim(str, cutset)
	}
	return result
}
