package main

import (
	"bufio"
	"fmt"
	"os"
)

type Scanner struct {
	Range   int
	Pos     int
	Reverse bool
}

type Depth int

func main() {
	scanners := map[Depth]*Scanner{}
	var maxDepth Depth
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var depth Depth
		var rng int
		_, err := fmt.Sscanf(s.Text(), "%d: %d", &depth, &rng)
		if err != nil {
			panic(err)
		}
		if depth > maxDepth {
			maxDepth = depth
		}
		scanners[depth] = &Scanner{Range: rng}
	}

	var severity int
	for d := Depth(0); d <= maxDepth; d++ {
		if s, ok := scanners[d]; ok && s.Pos == 0 {
			severity += int(d) * s.Range
		}
		for _, s := range scanners {
			if !s.Reverse {
				s.Pos++
			} else {
				s.Pos--
			}
			if s.Pos == 0 {
				s.Reverse = false
			} else if s.Pos == (s.Range - 1) {
				s.Reverse = true
			}
		}
	}
	fmt.Println("Severity:", severity)
}
