package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var freq int64
	for s.Scan() {
		v, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			log.Fatalf("failed to parse %q as an int: %v", s.Text(), err)
		}
		freq += v
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}
	fmt.Println(freq)
}
