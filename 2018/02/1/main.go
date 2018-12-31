package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func popCnt(s string) map[rune]int64 {
	pop := make(map[rune]int64)
	for _, r := range s {
		pop[r] += 1
	}
	return pop
}

func main() {
	var two, three int64
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var isTwo, isThree bool
		pop := popCnt(s.Text())
		for _, v := range pop {
			if v == 2 {
				isTwo = true
			}
			if v == 3 {
				isThree = true
			}
		}
		if isTwo {
			two += 1
		}
		if isThree {
			three += 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}
	fmt.Println(two * three)
}
