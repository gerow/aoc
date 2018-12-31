package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func ScanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' || data[i] == '\n' {
			return i + 1, data[:i], nil
		}
	}
	// If we're at EOF just return the last value.
	if atEOF {
		return len(data), data, bufio.ErrFinalToken
	}
	return 0, nil, nil
}

type State struct {
	Programs []byte
}

type MoveType int

const (
	Spin MoveType = iota
	Exchange
	Partner
)

type Move struct {
	Type                 MoveType
	SpinSize             int
	ExchangeA, ExchangeB int
	PartnerA, PartnerB   byte
}

func NewMove(m string) *Move {
	switch m[0] {
	case 's':
		var s int
		fmt.Sscanf(m, "s%d", &s)
		return &Move{Type: Spin, SpinSize: s}
	case 'x':
		var a, b int
		fmt.Sscanf(m, "x%d/%d", &a, &b)
		return &Move{Type: Exchange, ExchangeA: a, ExchangeB: b}
	case 'p':
		var a, b byte
		fmt.Sscanf(m, "p%c/%c", &a, &b)
		return &Move{Type: Partner, PartnerA: a, PartnerB: b}
	default:
		panic(m)
	}
}

func ParseMoves(r io.Reader) []Move {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanCommas)
	var out []Move
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		out = append(out, *NewMove(scanner.Text()))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return out
}

func New(n int) *State {
	var s State
	s.Programs = make([]byte, n)
	for i := range s.Programs {
		s.Programs[i] = 'a' + byte(i)
	}
	return &s
}

func (s *State) ApplyMoves(moves []Move) {
	for _, m := range moves {
		switch m.Type {
		case Spin:
			s.spin(m.SpinSize)
		case Exchange:
			s.exchange(m.ExchangeA, m.ExchangeB)
		case Partner:
			s.partner(m.PartnerA, m.PartnerB)
		}
	}
}

func (s *State) spin(n int) {
	p := make([]byte, len(s.Programs))
	for i := range p {
		p[i] = s.Programs[(len(p)-n+i)%len(p)]
	}
	s.Programs = p
}

func (s *State) exchange(a, b int) {
	s.Programs[a], s.Programs[b] = s.Programs[b], s.Programs[a]
}

func (s *State) partner(a, b byte) {
	var ai, bi int
	for i, v := range s.Programs {
		if v == a {
			ai = i
		}
		if v == b {
			bi = i
		}
	}
	s.exchange(ai, bi)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	moves := ParseMoves(os.Stdin)
	s := New(16)
	s.ApplyMoves(moves)
	fmt.Printf("end state: %q\n", string(s.Programs))

	s = New(16)
	seen := map[string]string{}
	for n := 0; n < 1000000000; n++ {
		before := string(s.Programs)
		if after, ok := seen[before]; ok {
			s.Programs = []byte(after)
			continue
		}
		s.ApplyMoves(moves)
		seen[before] = string(s.Programs)
	}
	fmt.Printf("end state after 1b: %q\n", string(s.Programs))
}
