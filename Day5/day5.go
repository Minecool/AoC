package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	buffer, _ := os.ReadFile("./input.txt")
	lowestValue := len(buffer)
	for i := 'a'; i <= '{'; i++ {
		data := string(buffer)
		prev := 0
		r := strings.NewReplacer(
			"Aa", "",
			"aA", "",
			"Bb", "",
			"bB", "",
			"Cc", "",
			"cC", "",
			"Dd", "",
			"dD", "",
			"Ee", "",
			"eE", "",
			"Ff", "",
			"fF", "",
			"Gg", "",
			"gG", "",
			"Hh", "",
			"hH", "",
			"Ii", "",
			"iI", "",
			"Jj", "",
			"jJ", "",
			"Kk", "",
			"kK", "",
			"Ll", "",
			"lL", "",
			"Mm", "",
			"mM", "",
			"Nn", "",
			"nN", "",
			"Oo", "",
			"oO", "",
			"Pp", "",
			"pP", "",
			"Qq", "",
			"qQ", "",
			"Rr", "",
			"rR", "",
			"Ss", "",
			"sS", "",
			"Tt", "",
			"tT", "",
			"Uu", "",
			"uU", "",
			"Vv", "",
			"vV", "",
			"Ww", "",
			"wW", "",
			"Xx", "",
			"xX", "",
			"Yy", "",
			"yY", "",
			"Zz", "",
			"zZ", "",
			string(i), "",
			strings.ToUpper(string(i)), "",
		)
	repeat:
		data = r.Replace(data)
		if prev != len(data) {
			prev = len(data)
			goto repeat
		}
		if len(data) < lowestValue {
			lowestValue = len(data)
		}
		if i == '{' {
			fmt.Println("Part 1:", len(data))
			fmt.Println("Part 2:", lowestValue)
		}
	}
}
