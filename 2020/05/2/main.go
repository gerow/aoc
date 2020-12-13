package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var cToN = map[rune]int{
	'F': 0,
	'B': 1,
	'L': 0,
	'R': 1,
}

func parse(s string) int {
	// looks like we can cheat pretty hard and just treat the ticket as a binary number with a simple mapping:
	// F => 0, B => 1, L => 0, R => 1
	// Of course, no checks that the ticket is *actually* valid.
	var id int
	for _, r := range s {
		d := cToN[r]
		id <<= 1
		id |= d
	}

	return id
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var ids []int
	for s.Scan() {
		ids = append(ids, parse(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// sort the list and look for the first gap
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i]-1 != ids[i-1] {
			fmt.Println(ids[i] - 1)
			os.Exit(0)
		}
	}

	log.Fatal("failed to find gap!")
}
