package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Pos struct {
	X, Y int
}

type Dir int

const (
	NE Dir = iota
	N
	NW
	SW
	S
	SE
)

var StringToDir = map[string]Dir{
	"ne": NE,
	"n":  N,
	"nw": NW,
	"sw": SW,
	"s":  S,
	"se": SE,
}

func (p Pos) Move(d Dir) Pos {
	switch d {
	case NE:
		return Pos{X: p.X + 1, Y: p.Y + 1}
	case N:
		return Pos{X: p.X, Y: p.Y + 2}
	case NW:
		return Pos{X: p.X - 1, Y: p.Y + 1}
	case SW:
		return Pos{X: p.X - 1, Y: p.Y - 1}
	case S:
		return Pos{X: p.X, Y: p.Y - 2}
	case SE:
		return Pos{X: p.X + 1, Y: p.Y - 1}
	}
	panic("invalid dir")
}

func (p Pos) Dist() int {
	// if x > y then x; else ceil(y/2)
	if p.X > p.Y {
		return p.X
	}
	d := p.Y / 2
	if p.Y%2 > 0 {
		d++
	}
	return d
}

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	ins := string(in)
	var dirs []Dir
	for _, d := range strings.Split(ins, ",") {
		dirs = append(dirs, StringToDir[d])
	}
	var p Pos
	var largest int
	for _, d := range dirs {
		p = p.Move(d)
		dist := p.Dist()
		if dist > largest {
			largest = dist
		}
	}
	fmt.Println(p)
	fmt.Println("Distance", p.Dist())
	fmt.Println("Largest distance", largest)
}
