package twelve

import (
	"fmt"
	"strings"
)

var rank = [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

var firstV = "On the %s day of Christmas my true love gave to me: "

var verses = [12]string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

//Song outputs the entire song
func Song() string {
	var sb strings.Builder
	for i := 1; i < 13; i++ {
		sb.WriteString(Verse(i) + "\n")
	}
	return sb.String()
}

//Verse outputs the song by line
func Verse(input int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(firstV, rank[input-1]))

	for i := 0; i < input-1; i++ {
		sb.WriteString(verses[input-i-1])
		if i < input-2 {
			sb.WriteString(", ")
		} else {
			sb.WriteString(", and ")
		}
	}
	sb.WriteString(verses[0] + ".")
	return sb.String()
}
