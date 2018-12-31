package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Claim struct {
	ID          int
	Start, Size Vec2
}

type Vec2 struct {
	X, Y int
}

func parse(s string) *Claim {
	var c Claim
	if _, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &c.ID, &c.Start.X, &c.Start.Y, &c.Size.X, &c.Size.Y); err != nil {
		panic(fmt.Sprintf("failed to parse %q: %v", c, err))
	}
	return &c
}

func overlaps(a, b *Claim) bool {
	// we can do this better, but just copying the multiclaim code from
	// part 1 is a quick and dirty solution
	claimed := make(map[Vec2]bool)

	// start by adding a
	for x := a.Start.X; x < a.Start.X+a.Size.X; x++ {
		for y := a.Start.Y; y < a.Start.Y+a.Size.Y; y++ {
			claimed[Vec2{x, y}] = true
		}
	}

	// now check every entry for b, if it's been covered by a already we
	// have overlap
	for x := b.Start.X; x < b.Start.X+b.Size.X; x++ {
		for y := b.Start.Y; y < b.Start.Y+b.Size.Y; y++ {
			if claimed[Vec2{x, y}] {
				return true
			}
		}
	}

	return false
}

func main() {
	var claims []*Claim
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		claims = append(claims, parse(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	for _, a := range claims {
		var failed bool
		for _, b := range claims {
			// a claim wil always overlap with itself
			if b.ID == a.ID {
				continue
			}
			if overlaps(a, b) {
				log.Printf("%d has overlap\n", a.ID)
				failed = true
				break
			}
		}
		if !failed {
			fmt.Println(a.ID)
			os.Exit(0)
		}
	}
	log.Fatalf("failed to find claim with no overlap")
}
