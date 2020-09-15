/*
isogram package is used to determins if a word or phrase hase any repeating letters, if it does not it is an isogram
*/

package isogram

import "unicode"

// IsIsogram determins if a word or phrase is an isogram
func IsIsogram(s string) bool {
	m := make(map[rune]bool)
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		r = unicode.ToLower(r)
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}
