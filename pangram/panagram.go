package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	res := true
	k := string(x)
	j := len(k)
	for _, i := range k {

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

func main(){
	b: = isPalindrome("-121")
fmt.Println(b)
}
