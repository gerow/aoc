package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var changes []int64
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatalf("failed to parse %q as an int: %v", s.Text(), err)
		}
		changes = append(changes, v)
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	var freq int64
	reached := make(map[int64]bool)
	reached[0] = true
	for {
		for _, v := range changes {
			freq += v
			if reached[freq] {
				fmt.Println(freq)
				os.Exit(0)
			}
			reached[freq] = true
		}
	}
}
