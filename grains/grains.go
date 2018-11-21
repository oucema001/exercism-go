package grains

import (
	"errors"
)

//Square calculates the number of grains on a given chess square
func Square(input int) (uint64, error) {
	var result int
	result = 1
	if input <= 0 || input > 64 {
		return 0, errors.New("case should be superior to 0 and less than 64")
	}
	result = 1 << (uint(input) - 1)
	return uint64(result), nil
}

//Total returns the sum of the grains on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
