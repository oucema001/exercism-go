package raindrops

import (
	"strconv"
)

//Convert : convererts number to string based on factors
func Convert(number int) string {
	drops := make([]string, 8, 8)
	drops[3] = "Pling"
	drops[5] = "Plang"
	drops[7] = "Plong"

	result := ""

	for i, v := range drops {
		if i != 0 && number%i == 0 {
			result += v
		}
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
