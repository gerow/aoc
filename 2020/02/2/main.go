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
	pos1, pos2 int
}

func newEntry(line string) (*entry, error) {
	var e entry
	var r string
	_, err := fmt.Sscanf(line, "%d-%d %1s: %s", &e.policy.pos1, &e.policy.pos2, &r, &e.password)
	if err != nil {
		return nil, err
	}
	e.policy.rune, _ = utf8.DecodeRuneInString(r)

	return &e, nil
}

func (e *entry) valid() bool {
	var p1, p2 bool
	// since we bothered to handle runes properly we do have to iterate through the password
	for i, r := range e.password {
		if r != e.policy.rune {
			continue
		}
		switch i+1 {
		case e.policy.pos1:
			p1 = true
		case e.policy.pos2:
			p2 = true
		default:
			continue
		}
	}
	// xor baybee
	return (p1 || p2) && !(p1 && p2)
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
