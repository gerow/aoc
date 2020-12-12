package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

type entry struct {
	password string
	policy   policy
}

type policy struct {
	rune     rune
	min, max int
}

func newEntry(line string) (*entry, error) {
	var e entry
	var r string
	_, err := fmt.Sscanf(line, "%d-%d %1s: %s", &e.policy.min, &e.policy.max, &r, &e.password)
	if err != nil {
		return nil, err
	}
	e.policy.rune, _ = utf8.DecodeRuneInString(r)

	return &e, nil
}

func (e *entry) valid() bool {
	var n int
	for _, r := range e.password {
		if r == e.policy.rune {
			n += 1
		}
	}
	return n >= e.policy.min && n <= e.policy.max
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var valid int
	for s.Scan() {
		e, err := newEntry(s.Text())
		if err != nil {
			log.Fatalf("failed to parse entry line %s: %v", s.Text(), err)
		}
		if e.valid() {
			valid += 1
		}
	}

	fmt.Println(valid)
}
