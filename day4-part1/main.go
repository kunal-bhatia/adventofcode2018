package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	sort.Strings(lines)

	var guard, asleep int
	C := make(map[int]int)
	CM := map[int]map[int]int{}

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
				if CM[guard] == nil {
					CM[guard] = map[int]int{}
				}
				CM[guard][i]++
			}
		}

	}
	bestGuard := findBestGuard(C)

	bestMin := findBestGuard(CM[bestGuard])
	fmt.Println(bestGuard * bestMin)

}

func parseTime(line string) (int, error) {
	t, error := strconv.Atoi(strings.Split(line[1:strings.IndexByte(line, ']')], ":")[1])
	return t, error
}

func findBestGuard(m map[int]int) int {
	best := 0
	for k, v := range m {
		if best == 0 || v > m[best] {
			best = k
		}
	}
	return best
}
