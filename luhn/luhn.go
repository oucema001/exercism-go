package luhn

import (
	"strconv"
)

//Valid: returns true if string valid
func Valid(input string) bool {
	if len(input) < 2 {
		return false
	}
	double := false
	//res := make([]int, 100)
	res2 := 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == ' ' {
			if len(input) == 2 {
				return false
			}
			continue
		}
		c, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return false
		}
		if double {
			c = c * 2
		}
		if c > 9 {
			c = c - 9
		}
		res2 += c
		double = !double
	}
	return res2%10 == 0
}
