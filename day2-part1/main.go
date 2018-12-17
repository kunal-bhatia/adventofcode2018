package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	characters := make([]int, 26)

	var duplicates, triplets int
	for s.Scan() {
		fillWithZeros(characters)

		for i := 0; i < len(s.Text()); i++ {
			c := string(s.Text()[i])
			characters[strings.Index(alphabet, c)]++
		}

		var twosFound, threesFound bool

		for _, num := range characters {
			if num == 2 {
				twosFound = true
			}

			if num == 3 {
				threesFound = true
			}
		}
		if twosFound {
			duplicates++
		}

		if threesFound {
			triplets++
		}
	}
	fmt.Println(duplicates * triplets)
}

func fillWithZeros(chars []int) *[]int {
	for i := range chars {
		chars[i] = 0
	}
	return &chars
}
