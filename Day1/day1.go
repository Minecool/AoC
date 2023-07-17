package main;
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"log"
)

func main() {

	num := 0;
	var data []string;
	seenBefore := make(map[int]bool);
	partOneDone := false;

	buffer, err := os.ReadFile("./input.txt");
	if err != nil {
		log.Fatal(err);
	}
	data = strings.Split(string(buffer), "\r\n");
	solve:
	for i:=0; i<len(data);i++ {
		if data[i] == "" {
			continue;
		}
		intData, err := strconv.Atoi(data[i]);
		if err != nil {
			log.Fatal(err);
		}
		num += intData;
		if seenBefore[num] == true {
			fmt.Println("Part 2:", num);
			return;
		}
		seenBefore[num] = true;
	}
	if !partOneDone {
		fmt.Println("Part 1:", num);
		partOneDone = true;
	}
	goto solve;
}
