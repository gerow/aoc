package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type xmas struct {
	prev []int
}

func (x *xmas) valid(n int) bool {
	if len(x.prev) != 25 {
		panic(fmt.Sprintf("somehow x.prev ended up being length %d", len(x.prev)))
	}

	// I think we can do this way *way* faster (but try it slow for now).
	for i := 0; i < len(x.prev)-1; i++ {
		for j := i + 1; j < len(x.prev); j++ {
			if x.prev[i]+x.prev[j] == n {
				return true
			}
		}
	}

	return false
}

func (x *xmas) extend(n int) {
	x.prev = append(x.prev[1:], n)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var x xmas
	for i := 0; i < 25; i++ {
		if !s.Scan() {
			if s.Err() != nil {
				log.Fatal(s.Err())
			}
			log.Fatalf("failed to read full preamble; read %d of 25 numbers", i)
		}
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to read %q as a number on line %d: %v", s.Text(), i+1, err)
		}
		x.prev = append(x.prev, n)
	}

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to read %q as a number: %v", s.Text(), err)
		}
		if !x.valid(n) {
			fmt.Println(n)
			os.Exit(0)
		}
		x.extend(n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	log.Fatal("failed to find an invalid number")
}
