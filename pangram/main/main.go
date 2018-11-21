package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	res := true
	k := string(x)
	j := len(k) - 1
	for i := 0; i < (len(k)/2)+1; i++ {

		if k[i] == k[j] {
			res = true
			i++
			j--
		} else {
			return false
		}
	}
	return res
}

func main() {
	var boo bool
	boo = isPalindrome(212)
	fmt.Println(boo)
}
