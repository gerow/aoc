package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	yes := make(map[rune]bool)
	var sum int
	for s.Scan() {
		if s.Text() == "" {
			// reached the end of a group. Count, reset, continue.
			sum += len(yes)
			yes = make(map[rune]bool)
			continue
		}
		for _, r := range s.Text() {
			yes[r] = true
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// need to count the last group
	sum += len(yes)

	fmt.Println(sum)
}
