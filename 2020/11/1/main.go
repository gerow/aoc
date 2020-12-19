package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

type vec2 [2]int

func (v1 vec2) add(v2 vec2) vec2 {
	return vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

type state struct {
	width, height int
	entries       []entry
}

type entry int

const (
	floor entry = iota
	empty
	occupied
)

func (s *state) load(r io.Reader) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		s.height++
		l := sc.Text()
		rc := utf8.RuneCountInString(l)
		if s.width == 0 {
			s.width = rc
		} else if rc != s.width {
			return fmt.Errorf("found line with wrong width %d (expecting %d)", rc, s.width)
		}

		for _, r := range l {
			var ent entry
			switch r {
			case '.':
				ent = floor
			case 'L':
				ent = empty
			case '#':
				ent = occupied
			default:
				return fmt.Errorf("invalid entry %v", r)
			}
			s.entries = append(s.entries, ent)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *state) clone() *state {
	var s2 state
	s2 = *s
	s2.entries = make([]entry, len(s.entries))
	copy(s2.entries, s.entries)

	return &s2
}

func (s *state) adjOccupied(v vec2) int {
	var occ int
	for _, d := range []vec2{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	} {
		l := v.add(d)
		// limits check
		if l[0] < 0 || l[1] < 0 || l[0] >= s.width || l[1] >= s.height {
			continue
		}
		if s.at(l) == occupied {
			occ++
		}
	}

	return occ
}

func (s *state) step() (changed bool) {
	prev := s.clone()

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			v := vec2{x, y}
			e := prev.at(v)
			switch e {
			case floor:
				continue
			case empty:
				if prev.adjOccupied(v) == 0 {
					s.set(v, occupied)
					changed = true
				}
			case occupied:
				if prev.adjOccupied(v) >= 4 {
					s.set(v, empty)
					changed = true
				}
			}
		}
	}

	return changed
}

func (s *state) at(v vec2) entry {
	return s.entries[v[1]*s.width+v[0]]
}

func (s *state) set(v vec2, e entry) {
	s.entries[v[1]*s.width+v[0]] = e
}

func main() {
	var s state
	if err := s.load(os.Stdin); err != nil {
		log.Fatal(err)
	}
	for s.step() {
	}

	var occ int
	for _, e := range s.entries {
		if e == occupied {
			occ++
		}
	}

	fmt.Println(occ)
}
