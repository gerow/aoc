package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type State int

const (
	Normal State = iota
	InGarbage
)

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	b := data[0]
	if b == '{' || b == '}' || b == ',' {
		return 1, []byte{b}, nil
	}
	if b == '<' {
		var cancel bool
		for i, d := range data {
			b = d
			if b == '>' && !cancel {
				return i + 1, data[:i+1], nil
			}
			if cancel {
				cancel = false
			} else if b == '!' {
				cancel = true
			}
		}
		// Need more data to finish garbage.
		return 0, nil, nil
	}
	// Ignore newline.
	if b == '\n' {
		return 1, nil, nil
	}
	return 0, nil, fmt.Errorf("failed to tokenize string beginning with %q", b)
}

func Score(r io.Reader) int {
	var depth, score int
	s := bufio.NewScanner(r)
	s.Split(ScanTokens)
	for s.Scan() {
		t := s.Text()
		if t[0] == '{' {
			depth++
			score += depth
		} else if t[0] == '}' {
			depth--
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return score
}

func main() {
	fmt.Println("Score:", Score(os.Stdin))
}
