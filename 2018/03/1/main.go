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

func main() {
	var claims []*Claim
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		claims = append(claims, parse(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	// now that we have the claims we just count up the number of times
	// each point is claimed by an elf.
	claimed := make(map[Vec2]int)
	for _, c := range claims {
		for x := c.Start.X; x < c.Start.X+c.Size.X; x++ {
			for y := c.Start.Y; y < c.Start.Y+c.Size.Y; y++ {
				claimed[Vec2{x, y}] += 1
			}
		}
	}

	// now just count up the points claimed by multiple elves
	var overlap int
	for _, n := range claimed {
		if n > 1 {
			overlap += 1
		}
	}
	fmt.Println(overlap)
}
