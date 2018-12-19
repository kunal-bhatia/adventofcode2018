package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := strings.TrimSpace(string(b))
	ok := true
	for ok {
		ok = false
		for i := 0; i < len(str)-1; i++ {
			current, next := str[i], str[i+1]
			if current == next+32 || current == next-32 {
				str = str[:i] + str[i+2:]
				ok = true
				break
			}
		}
	}
	fmt.Println(len(str))
}
