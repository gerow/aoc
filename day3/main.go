package main

import "fmt"

type Point struct {
	X int
	Y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func shell(pt Point) int {
	return max(abs(pt.X), abs(pt.Y))
}

func next(pt Point) Point {
	s := shell(pt)
	// We're at the bottom right of this shell, move on to the next one.
	if pt.X == s && pt.Y == -s {
		return Point{pt.X + 1, pt.Y}
	}
	// We're on the right with more room going up
	if pt.X == s && pt.Y < s {
		return Point{pt.X, pt.Y + 1}
	}
	// We're on the top right corner, start working left
	if pt.X == s && pt.Y == s {
		return Point{pt.X - 1, pt.Y}
	}
	// We're on the top with more room going left
	if pt.Y == s && pt.X > -s {
		return Point{pt.X - 1, pt.Y}
	}
	// We're on the top left corner, start working down
	if pt.X == -s && pt.Y == s {
		return Point{pt.X, pt.Y - 1}
	}
	// We're on the left with more room going down
	if pt.X == -s && pt.Y > -s {
		return Point{pt.X, pt.Y - 1}
	}
	// We're on the bottom left corner, work right
	if pt.X == -s && pt.Y == -s {
		return Point{pt.X + 1, pt.Y}
	}
	// We're on the bottom with more room going right
	if pt.Y == -s && pt.X < s {
		return Point{pt.X + 1, pt.Y}
	}
	panic("this should never happen")
}

func sum(pt Point, ents map[Point]int) int {
	s := 0

	for _, tp := range []Point{
		{pt.X + 1, pt.Y},
		{pt.X + 1, pt.Y + 1},
		{pt.X, pt.Y + 1},
		{pt.X - 1, pt.Y + 1},
		{pt.X - 1, pt.Y},
		{pt.X - 1, pt.Y - 1},
		{pt.X, pt.Y - 1},
		{pt.X + 1, pt.Y - 1},
	} {
		if v, ok := ents[tp]; ok {
			s += v
		}
	}
	return s
}

func main() {
	n := 368078

	pt := Point{0, 0}
	ents := map[Point]int{}
	ents[pt] = 1
	pt = next(pt)
	for ; ; pt = next(pt) {
		s := sum(pt, ents)
		if s > n {
			fmt.Println("sum:", s)
			break
		}
		ents[pt] = s
	}
}
