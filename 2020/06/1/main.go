package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	yes := make(map[rune]bool)
	var sum int
	for s.Scan() {
		if s.Text() == "" {
			// reached the end of a group. Count, reset, continue.
			for range yes {
				sum++
			}
			yes = make(map[rune]bool)
		}
		for _, r := range s.Text() {
			yes[r] = true
		}
	}
	// need to count the last group
	for range yes {
		sum++
	}

	fmt.Println(sum)
}
