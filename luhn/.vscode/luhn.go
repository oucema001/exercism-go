package main

import (
	"fmt"
)

//Valid: returns true if string valid
func Valid(input string) bool {

	if len(input) < 2 {
		return false
	}
	double := false
	res := make([]int, len(input))
	for i := len(input) - 1; i >= 0; i-- {

		c := int(input[i] - '0')
		if double {
			c = c * 2
		}
		if c > 9 {
			c = c - 9
		}
		res[i] = c
		double = !double
	}
	res2 := 0
	for i := range res {
		res2 += res[i]
	}
	return res2%10 == 0

}

func main() {
	fmt.Println(Valid("4539 1488 0343 6467"))
}
