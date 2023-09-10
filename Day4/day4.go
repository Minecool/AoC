package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type parsedLog struct {
	minute      int
	time        int
	id          int
	instruction string
}

func formatLog(logs []string) []parsedLog {
	var result []parsedLog
	r := strings.NewReplacer(
		"-", "",
		" ", "",
		":", "",
	)
	for _, log := range logs {
		var formattedLog parsedLog
		splitBySpace := strings.Split(log, " ")
		formattedLog.minute, _ = strconv.Atoi(log[15:17])
		formattedLog.time, _ = strconv.Atoi(r.Replace(log[1:17]))
		formattedLog.instruction = splitBySpace[len(splitBySpace)-2] + " " + splitBySpace[len(splitBySpace)-1]
		if formattedLog.instruction == "begins shift" {
			formattedLog.id, _ = strconv.Atoi(strings.Split(strings.Split(log, "#")[1], " ")[0])
		} else {
			formattedLog.id = -1
		}
		result = append(result, formattedLog)
	}
	return result
}

func sortLog(log []parsedLog) []parsedLog {
	sort.Slice(log, func(i, j int) bool {
		return log[i].time < log[j].time
	})
	return log
}

func getTimeAsleep(logs []parsedLog) map[string]int {
	guards := make(map[string]int)
	currentGuard := -1
	asleep := 0
	for _, log := range logs {
		if log.id == -1 {
			log.id = currentGuard
		} else {
			currentGuard = log.id
		}
		if log.instruction == "falls asleep" {
			asleep = log.minute
		}
		if log.instruction == "wakes up" {
			for i := asleep; i < log.minute; i++ {
				guards[strconv.Itoa(log.id)+"_"+strconv.Itoa(i)]++
			}
		}
	}
	return guards
}

func solve(logs []parsedLog) {
	guards := getTimeAsleep(logs)
	guardsWithTotalMinutesAsleep := make(map[string]int)

	part1guard := 0
	part1minute := 0

	part2guard := 0
	part2minute := 0

	maxValue := 0

	for guard, amount := range guards {
		guardsWithTotalMinutesAsleep[strings.Split(guard, "_")[0]] += amount
		if amount > maxValue {
			maxValue = amount
			part2minute, _ = strconv.Atoi(strings.Split(guard, "_")[1])
			part2guard, _ = strconv.Atoi(strings.Split(guard, "_")[0])
		}
	}
	maxValue = 0

	for guard, minutesAsleep := range guardsWithTotalMinutesAsleep {
		if minutesAsleep > maxValue {
			maxValue = minutesAsleep
			part1guard, _ = strconv.Atoi(guard)
		}
	}
	maxValue = 0

	for i := 0; i < 60; i++ {
		minute := guards[strconv.Itoa(part1guard)+"_"+strconv.Itoa(i)]
		if minute > maxValue {
			maxValue = minute
			part1minute = i
		}
	}
	fmt.Println("Part 1:", part1guard*part1minute)
	fmt.Println("Part 2:", part2guard*part2minute)
}

func main() {
	buffer, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(buffer), "\r\n")

	var formattedLog []parsedLog = formatLog(lines)
	var sortedLog []parsedLog = sortLog(formattedLog)
	solve(sortedLog)
}
