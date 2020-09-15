// Package grains Calculates the number of grains of wheat on a chessboard given that the number on each square doubles
package grains

import (
	"fmt"
)

// Square returns the number of grains on the specified square of the chessboard, or an error if the square is invalid.
func Square(input int) (actualVal uint64, actualErr error) {
	if input < 1 || input > 64 {
		return 0, fmt.Errorf("input(%d) should be between 1 and 64", input)
	}
	return 1 << (input - 1), nil

}

// Total returns the total amount of grain
func Total() uint64 {
	// equivalent to  2**64-1
	return 1<<64 - 1
}
