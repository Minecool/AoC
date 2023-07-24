package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var data []string

	buffer, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = strings.Split(string(buffer), "\r\n")

	var twinsCount int
	var tripletsCount int
	var part2done bool
	for i := 0; i < len(data); i++ {
		var twins bool
		var triplets bool
		countChars := make(map[byte]int)
		var line string = data[i]

		for j := 0; j < len(line); j++ {
			countChars[line[j]] += 1
			var points int
			for k := 0; k < len(data); k++ {
				var comparedLine string = data[k]
				if !part2done {
					for l := 0; l < len(line); l++ {
						if line[l] == comparedLine[l] {
							points++
						}
					}
					if points == 25 {
						fmt.Print("Part 2: ")
						for o := 0; o < len(line); o++ {
							if line[o] == comparedLine[o] {
								fmt.Print(string(line[o]))
							}
						}
						fmt.Print("\n")
						part2done = true
					}
					points = 0
				}
			}
		}

		for _, value := range countChars {

			if value == 2 {
				twins = true
			}
			if value == 3 {
				triplets = true
			}
		}
		if twins {
			twinsCount += 1
		}
		if triplets {
			tripletsCount += 1
		}
	}
	fmt.Println("Part 1:", twinsCount*tripletsCount)
}
