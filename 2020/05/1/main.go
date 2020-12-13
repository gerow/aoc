package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var highest int
	for s.Scan() {
		v := parse(s.Text())
		if v > highest {
			highest = v
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(highest)
}
