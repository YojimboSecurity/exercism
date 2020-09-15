package luhn

import (
	"strconv"
	"strings"
)

func luhn(n int) int {
	n *= 2
	if n >= 10 {
		n -= 9
	}
	return n
}

// Valid returns true if luhnString is a valid Luhn number
func Valid(luhnString string) bool {
	// remove spaces
	luhnString = strings.ReplaceAll(luhnString, " ", "")
	// check length is greater than 1
	luhnLength := len(luhnString)
	if luhnLength <= 1 {
		return false
	}
	// check if even or odd
	even := luhnLength%2 == 0
	source := strings.Split(luhnString, "")
	checksum := 0
	// loop over slice
	for i := len(source) - 1; i > -1; i-- {
		// check that it is an integer
		t, err := strconv.ParseInt(source[i], 10, 8)
		if err != nil {
			return false
		}
		n := int(t)
		if even == (i%2 == 0) {
			n = luhn(n)
		}

		checksum += n
	}
	return checksum%10 == 0
}
