package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	id, x, y, w, h int
}

type xy struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var claims []claim
	for s.Scan() {
		str := s.Text()
		r, _ := regexp.Compile("(#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+))")
		res := r.FindAllStringSubmatch(str, -1)
		for i := range res {
			c, _ := strconv.Atoi(res[i][2])
			x, _ := strconv.Atoi(res[i][3])
			y, _ := strconv.Atoi(res[i][4])
			w, _ := strconv.Atoi(res[i][5])
			h, _ := strconv.Atoi(res[i][6])
			//fmt.Printf("claimid=%s, x=%s, y=%s, w=%s, h=%s\n", res[i][2], res[i][3], res[i][4], res[i][5], res[i][6])
			claims = append(claims, claim{c, x, y, w, h})
		}
	}

	dict := make(map[xy]int)

	for _, claim := range claims {
		for i := 0; i < claim.w; i++ {
			for j := 0; j < claim.h; j++ {
				dict[xy{claim.x + i, claim.y + j}]++
			}
		}
	}

	ans := 0
	for _, v := range dict {
		if v > 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
