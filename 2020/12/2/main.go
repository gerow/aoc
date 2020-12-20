package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type vec2 [2]int

func (v1 vec2) add(v2 vec2) vec2 {
	return vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v vec2) scale(s int) vec2 {
	return vec2{v[0] * s, v[1] * s}
}

// Weird conventions here, but m just indicates the number of times we rotate left by 90 degrees.
func (v vec2) rotate(m int) vec2 {
	switch m {
	case 1:
		return vec2{-v[1], v[0]}
	case 2:
		return vec2{-v[0], -v[1]}
	case 3:
		return vec2{v[1], -v[0]}
	default:
		panic(fmt.Sprintf("invalid rotation %d", m))
	}
}
var (
	north = vec2{0, 1}
	south = vec2{0, -1}
	east  = vec2{1, 0}
	west  = vec2{-1, 0}
)

type movementType int

const (
	moveWaypoint movementType = iota
	rotateWaypoint
	forward
)

type movement struct {
	typ movementType

	dir vec2
	mag int
}

func parseMovement(s string) (*movement, error) {
	var m movement
	var dir string

	_, err := fmt.Sscanf(s, "%1s%d", &dir, &m.mag)
	if err != nil {
		return nil, fmt.Errorf("failed to parse movement %q: %w", s, err)
	}

	switch dir {
	case "N":
		m.typ = moveWaypoint
		m.dir = north
	case "S":
		m.typ = moveWaypoint
		m.dir = south
	case "E":
		m.typ = moveWaypoint
		m.dir = east
	case "W":
		m.typ = moveWaypoint
		m.dir = west
	case "L":
		m.typ = rotateWaypoint
		m.mag = m.mag / 90
	case "R":
		m.typ = rotateWaypoint
		m.mag = m.mag/90
		// turn right turns into left turns
		m.mag = 4 - m.mag
	case "F":
		m.typ = forward
	default:
		return nil, fmt.Errorf("invalid direction %q", dir)
	}

	return &m, nil
}

type boat struct {
	waypoint, loc vec2
}

func (b *boat) apply(m *movement) {
	switch m.typ {
	case moveWaypoint:
		b.waypoint = b.waypoint.add(m.dir.scale(m.mag))
	case rotateWaypoint:
		b.waypoint = b.waypoint.rotate(m.mag)
	case forward:
		b.loc = b.loc.add(b.waypoint.scale(m.mag))
	default:
		panic(fmt.Sprintf("unknown movement type %v", m.typ))
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var b boat
	b.waypoint = vec2{10, 1}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		m, err := parseMovement(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		b.apply(m)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(abs(b.loc[0]) + abs(b.loc[1]))
}
