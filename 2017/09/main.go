package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	d := data[0]
	if d == '{' || d == '}' || d == ',' {
		return 1, []byte{d}, nil
	}
	if d == '<' {
		var cancel bool
		var out []byte
		for i, d := range data {
			if d == '>' && !cancel {
				out = append(out, d)
				return i + 1, out, nil
			}
			if cancel {
				cancel = false
			} else if d == '!' {
				cancel = true
			} else {
				out = append(out, d)
			}
		}
		// Need more data to finish garbage.
		return 0, nil, nil
	}
	// Ignore newline.
	if d == '\n' {
		return 1, nil, nil
	}
	return 0, nil, fmt.Errorf("failed to tokenize string beginning with %q", d)
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

func GarbageCount(r io.Reader) int {
	var garbage int
	s := bufio.NewScanner(r)
	s.Split(ScanTokens)
	for s.Scan() {
		t := s.Text()
		if t[0] == '<' {
			garbage += len(t) - 2
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return garbage
}

func main() {
	var b bytes.Buffer
	t := io.TeeReader(os.Stdin, &b)
	fmt.Println("Score:", Score(t))
	fmt.Println("GarbageCount:", GarbageCount(&b))
}
