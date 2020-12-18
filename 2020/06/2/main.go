package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func intersect(a, b map[rune]bool) map[rune]bool {
	var s, l map[rune]bool
	if len(a) < len(b) {
		s, l = a, b
	} else {
		s, l = b, a
	}

	o := make(map[rune]bool)
	for k := range s {
		if l[k] {
			o[k] = true
		}
	}

	return o
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var yesAll map[rune]bool
	var sum int
	first := true
	// this got way too hairy way to quick
	for s.Scan() {
		if s.Text() == "" {
			// reached the end of a group. Count, reset, and continue.
			sum += len(yesAll)
			first = true
			continue
		}
		thisYes := make(map[rune]bool)
		for _, r := range s.Text() {
			thisYes[r] = true
		}
		if first {
			yesAll = thisYes
			first = false
		} else {
			yesAll = intersect(yesAll, thisYes)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// need to count the last group
	sum += len(yesAll)

	fmt.Println(sum)
}
