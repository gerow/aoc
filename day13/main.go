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
type Scanners map[Depth]*Scanner

func (ss Scanners) Step() {
	for _, s := range ss {
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

func (ss Scanners) Reset() {
	for _, s := range ss {
		s.Pos = 0
		s.Reverse = false
	}
}

func main() {
	scanners := Scanners{}
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
		scanners.Step()
	}
	fmt.Println("Severity:", severity)

	var delay int
	for ; ; delay++ {
		scanners.Reset()
		for n := 0; n < delay; n++ {
			scanners.Step()
		}
		caught := false
		for d := Depth(0); d <= maxDepth; d++ {
			if s, ok := scanners[d]; ok && s.Pos == 0 {
				caught = true
				break
			}
			scanners.Step()
		}
		if !caught {
			break
		}
	}
	fmt.Println("Minimum delay:", delay)
}
