package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type square int

const (
	empty square = iota
	tree
)

type grid struct {
	squares       []square
	width, height int
}

type vec struct {
	x, y int
}

func (v vec) add(v2 vec) vec {
	return vec{v.x + v2.x, v.y + v2.y}
}

func (g *grid) at(v vec) square {
	// wrap in the x direction
	v.x %= g.width
	if v.y > g.height {
		panic(fmt.Sprintf("index y %d beyond range", v.y))
	}
	return g.squares[v.y*g.width+v.x]
}

func newGrid(r io.Reader) (*grid, error) {
	s := bufio.NewScanner(r)
	first := true
	var g grid
	for s.Scan() {
		l := 0
		for _, r := range s.Text() {
			l += 1
			if first {
				g.width += 1
			}
			var sq square
			switch r {
			case '.':
				sq = empty
			case '#':
				sq = tree
			default:
				return nil, fmt.Errorf("got invalid square with rune %s", string(r))
			}
			g.squares = append(g.squares, sq)
		}
		g.height += 1
		first = false
		// ensure this line matches our width
		if l != g.width {
			return nil, fmt.Errorf("malformed grid, 1st line had width %d, %dth line had width %d", g.width, g.height, l)
		}
	}

	return &g, nil
}

func main() {
	g, err := newGrid(os.Stdin)
	if err != nil {
		log.Fatalf("failed to parse grid: %v", err)
	}

	cases := []vec{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	mul := 1
	for _, c := range cases {
		var loc vec
		var trees int
		for loc.y < g.height {
			if g.at(loc) == tree {
				trees += 1
			}
			loc = loc.add(c)
		}
		mul *= trees
	}

	fmt.Println(mul)
}
