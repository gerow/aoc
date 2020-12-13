package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

func newPassport(r io.Reader) (*passport, error) {
	var p passport

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		var cat, val string
		_, err := fmt.Sscanf(s.Text(), "%3s:%s", &cat, &val)
		if err != nil {
			return nil, fmt.Errorf("failed to scanf %q: %w", s.Text(), err)
		}
		var t *string
		switch cat {
		case "byr":
			t = &p.birthYear
		case "iyr":
			t = &p.issueYear
		case "eyr":
			t = &p.expirationYear
		case "hgt":
			t = &p.height
		case "hcl":
			t = &p.hairColor
		case "ecl":
			t = &p.eyeColor
		case "pid":
			t = &p.passportID
		case "cid":
			t = &p.countryID
		}
		if *t != "" {
			return nil, fmt.Errorf("field %v defined twice", cat)
		}
		*t = val
	}
	return &p, s.Err()
}

func (p *passport) valid() bool {
	return p.birthYear != "" &&
		p.issueYear != "" &&
		p.expirationYear != "" &&
		p.height != "" &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportID != ""
}

func scanDoubleLines(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(scanDoubleLines)

	var valid int
	for s.Scan() {
		p, err := newPassport(bytes.NewReader(s.Bytes()))
		if err != nil {
			log.Fatal(err)
		}
		if p.valid() {
			valid += 1
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(valid)
}
