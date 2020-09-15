/*
raindrops package provides a function to convert a number into a string that contains raindrop sounds corresponding to certain potential factors.
*/

package raindrops

import "strconv"

// Convert takes a number and returns a sound or the number as a string
func Convert(i int) string {
	var s string
	if i%3 == 0 {
		s += "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}
	if s != "" {
		return s
	}
	// convert the int into a string
	s = strconv.Itoa(i)
	return s
}
