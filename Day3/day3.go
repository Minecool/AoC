package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	id   string
	minX int
	minY int
	maxX int
	maxY int
}

func parseLine(line string) tile {
	id := strings.Split(line, " @")[0]                       // #1 @ 82,901: 26x12 -> #1
	area := strings.Split(line, "@ ")[1]                     // #1 @ 82,901: 26x12 -> 82,901: 26x12
	starting := strings.Split(area, ": ")[0]                 // 82,901: 26x12 -> 82,901
	lengths := strings.Split(area, ": ")[1]                  // 82,901: 26x12 -> 26x12
	minX, _ := strconv.Atoi(strings.Split(starting, ",")[0]) // 82,901 -> 82
	minY, _ := strconv.Atoi(strings.Split(starting, ",")[1]) // 82,901 -> 901
	xLen, _ := strconv.Atoi(strings.Split(lengths, "x")[0])  // 26x12 -> 26
	yLen, _ := strconv.Atoi(strings.Split(lengths, "x")[1])  // 26x12 -> 12

	var tile tile

	tile.id = id
	tile.minX = minX
	tile.minY = minY
	tile.maxX = minX + xLen
	tile.maxY = minY + yLen

	return tile
}

func main() {
	buffer, _ := os.ReadFile("./input.txt")
	rawData := string(buffer)
	lines := strings.Split(rawData, "\r\n")
	data := make([]tile, len(lines), len(lines)*2)
	for index, line := range lines {
		data[index] = parseLine(line)
	}
	seenBefore := make(map[string]int)
	for i := 0; i < len(data); i++ {
		for x := data[i].minX; x < data[i].maxX; x++ {
			for y := data[i].minY; y < data[i].maxY; y++ {
				loc := strconv.Itoa(x) + "," + strconv.Itoa(y)
				seenBefore[loc]++
			}
		}
		num := 0
		for j := 0; j < len(data); j++ {
			if j == i { // to skip when matching with itself
				continue
			}

			xOverlaps := data[i].minX <= data[j].maxX && data[i].maxX >= data[j].minX
			yOverlaps := data[i].minY <= data[j].maxY && data[i].maxY >= data[j].minY

			if xOverlaps && yOverlaps {
				num++
				break
			}

		}
		if num == 0 {
			fmt.Println("Part 2:", strings.Replace(data[i].id, "#", "", -1))
		}
	}
	part1 := 0
	for _, index := range seenBefore {
		if index > 1 {
			part1++
		}
	}
	fmt.Println("Part 1", part1)
}
