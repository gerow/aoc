package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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

func New(n int) *State {
	var s State
	s.Programs = make([]byte, n)
	for i := range s.Programs {
		s.Programs[i] = 'a' + byte(i)
	}
	return &s
}

func (s *State) Apply(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanCommas)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		switch t[0] {
		case 's':
			var n int
			fmt.Sscanf(t, "s%d", &n)
			s.spin(n)
		case 'x':
			var a, b int
			fmt.Sscanf(t, "x%d/%d", &a, &b)
			s.exchange(a, b)
		case 'p':
			var a, b byte
			fmt.Sscanf(t, "p%c/%c", &a, &b)
			s.partner(a, b)
		default:
			panic(t)
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
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

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	s := New(16)
	buf := bytes.NewReader(in)
	s.Apply(buf)
	fmt.Printf("end state: %q\n", string(s.Programs))

	s = New(16)
	for n := 0; n < 2000; n++ {
		buf.Reset(in)
		s.Apply(buf)
	}
	fmt.Printf("end state after 1b: %q\n", string(s.Programs))
}
