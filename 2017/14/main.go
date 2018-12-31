package main

import (
	"fmt"
	"math/bits"

	"github.com/gerow/aoc/day10/knot"
)

const totalRows = 128
const totalCols = 128

type Grid struct {
	Rows [totalRows][16]byte
}

type Point struct {
	Row, Col int
}

func New(name string) *Grid {
	var g Grid
	for r := 0; r < len(g.Rows); r++ {
		g.Rows[r] = knot.Sum([]byte(fmt.Sprintf("%s-%d", name, r)))
	}
	return &g
}

func (g *Grid) Used() int {
	used := 0
	for _, r := range g.Rows {
		for _, b := range r {
			used += bits.OnesCount8(b)
		}
	}
	return used
}

func (g *Grid) Get(row, col int) bool {
	return (0x80>>uint(col%8))&g.Rows[row][col/8] != 0
}

func (g *Grid) Groups() int {
	seen := map[Point]bool{}
	var q []Point
	var groups int
	for c := 0; c < totalCols; c++ {
		for r := 0; r < totalRows; r++ {
			if _, ok := seen[Point{r, c}]; ok {
				continue
			}
			if !g.Get(r, c) {
				continue
			}
			groups++
			// flood fill
			q = append(q, Point{r, c})
			for len(q) != 0 {
				c := q[0]
				q = q[1:]
				if _, ok := seen[c]; ok {
					continue
				}
				seen[c] = true
				if c.Row >= totalRows || c.Row < 0 || c.Col >= totalCols || c.Row < 0 {
					continue
				}
				if g.Get(c.Row, c.Col) {
					q = append(q,
						Point{c.Row + 1, c.Col},
						Point{c.Row - 1, c.Col},
						Point{c.Row, c.Col + 1},
						Point{c.Row, c.Col - 1})
				}
			}
		}
	}
	return groups
}

func main() {
	g := New("hxtvlmkl")

	fmt.Println("Used:", g.Used())
	fmt.Println("Groups:", g.Groups())
}
