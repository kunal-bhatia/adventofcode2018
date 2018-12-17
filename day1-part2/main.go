package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		var n int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)

	}
	m := map[int]bool{0: true}
	sum := 0
	for {
		for _, n := range numbers {
			//fmt.Printf("number is %d, sum is %d\n", n, sum)
			sum += n
			if m[sum] {
				fmt.Println(sum)
				return
			}
			m[sum] = true
		}
	}
}
