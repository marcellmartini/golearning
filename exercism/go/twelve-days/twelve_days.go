package twelve

import "fmt"

var lyric = [][]string{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves, and "},
	{"third", "three French Hens, "},
	{"fourth", "four Calling Birds, "},
	{"fifth", "five Gold Rings, "},
	{"sixth", "six Geese-a-Laying, "},
	{"seventh", "seven Swans-a-Swimming, "},
	{"eighth", "eight Maids-a-Milking, "},
	{"ninth", "nine Ladies Dancing, "},
	{"tenth", "ten Lords-a-Leaping, "},
	{"eleventh", "eleven Pipers Piping, "},
	{"twelfth", "twelve Drummers Drumming, "},
}

func Verse(i int) string {
	var line string
	var dayOut string
	for k, v := range lyric {
		if k < i {
			line = v[1] + line
			dayOut = v[0]
		}
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", dayOut, line)
}

func Song() string {
	var song string
	for i := 0; i < 12; i++ {
		song += Verse(i+1) + "\n"
	}

	return song[:len(song)-1]
}
