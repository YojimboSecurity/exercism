/*
Hamming package is used to calculate the Hamming distange between two strings.
The hamming distance is the differences between these two strings.

For example:

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^

This has the haming distance of 7.
*/

package hamming

import (
	"errors"
)

// Distance function calculates the hamming distance for sequences of strings a and
// b and will return the hamming distance as an integer. An error will be returned
// if the sequences are not of equal length.
func Distance(a, b string) (int, error) {
	// Check the length of the two strings returning error if different
	if len(a) != len(b) {
		return 0, errors.New("string sequence is of different lengths")
	}
	
	// Check each slices position for a match
	hamming, aa, bb := 0, []rune(a), []rune(b)
	for i, c := range aa {
		if c != bb[i] {
			hamming ++
		}
	}

	return hamming, nil
}
