package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type guardTime struct {
	id, time int
}

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	sort.Strings(lines)

	var guard, asleep int
	C := make(map[int]int)
	CM := make(map[guardTime]int)

	for _, line := range lines {
		t, err := parseTime(line)
		if err != nil {
			log.Fatal(err)
		}
		fields := strings.Fields(line[strings.Index(line, "]")+2:])
		if strings.Contains(line, "begins shift") {
			guard, _ = strconv.Atoi(fields[1][1:])
			asleep = 0
		} else if strings.Contains(line, "falls asleep") {
			asleep = t
		} else if strings.Contains(line, "wakes up") {
			for i := asleep; i < t; i++ {
				C[guard]++
				CM[guardTime{guard, i}]++
			}
		}

	}
	best, bestGuard, bestMin := 0, 0, 0

	for k, v := range CM {
		if v > best {
			best, bestGuard, bestMin = v, k.id, k.time
		}
	}
	fmt.Println(bestGuard * bestMin)

}

func parseTime(line string) (int, error) {
	t, error := strconv.Atoi(strings.Split(line[1:strings.IndexByte(line, ']')], ":")[1])
	return t, error
}
