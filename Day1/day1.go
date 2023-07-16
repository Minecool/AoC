package main;
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func solve() {
	for i:=0; i<len(splitData);i++ {
		intSplitData, err := strconv.Atoi(splitData[i]); // no idea what Atoi means
		if err != nil {
			fmt.Println(err);
			return;
		}
		num = num + intSplitData;
		for j:=0;j<len(seenBefore);j++ {
			if num == seenBefore[j] {
				fmt.Println("Part 2:", num);
				return;
			}
		}
		seenBefore = append(seenBefore, num);
	}
	if !partOneDone {
	fmt.Println("Part 1:", num);
	partOneDone = true;
	}
	solve();
}

var num int = 0;
var splitData []string;
var seenBefore []int;
var partOneDone bool = false;

func main() {
	buffer, err := os.ReadFile("./input.txt");
	if err != nil {
		fmt.Println(err);
		return;
	}
	var data string = string(buffer);
	splitData = strings.Split(data, "\r\n");
	solve()
}
