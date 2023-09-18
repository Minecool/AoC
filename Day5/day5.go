package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	buffer, _ := os.ReadFile("./input.txt")
	part1 := len(buffer)
	part2 := len(buffer)
	for char := 'a'; char <= '{'; char++ {
		// data := buffer // the sole reason my first solution wouldn't work.. everything after the first loop had a length of 50k..
		data := make([]byte, len(buffer))
		copy(data, buffer)
		for i := 0; i < len(data)-1; i++ {
			if data[i] == byte(char) || data[i] == byte(char-0x20) { // ended up being cleaner than strings.ToUpper / thanks hope
				data = append(data[:i], data[i+1:]...) // thanks hope
				if i == 0 {                            // doesn't feel too good running this check every time
					i++
				}
				i--
			}
			if data[i]-data[i+1] == 0x20 || data[i+1]-data[i] == 0x20 { // could probably make this its own function but meh
				data = append(data[:i], data[i+2:]...)
				if i == 0 {
					i++
				}
				i -= 2
			}
		}
		if len(data) < part2 {
			part2 = len(data)
		}
		if char == '{' {
			part1 = len(data)
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
