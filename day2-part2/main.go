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
	var ids []string
	for s.Scan() {
		ids = append(ids, s.Text())
	}

	for _, x := range ids {
		for _, y := range ids {
			diff := 0
			for i := 0; i < len(x); i++ {
				if x[i] != y[i] {
					diff++
				}
			}
			if diff == 1 {
				ans := make([]string, 0)
				for i := 0; i < len(x); i++ {
					if x[i] == y[i] {
						ans = append(ans, string(x[i]))
					}
				}
				fmt.Println(strings.Join(ans, ""))
				os.Exit(0)
			}
		}
	}
}
