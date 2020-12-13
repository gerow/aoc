package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	if !(p.birthYear != "" &&
		p.issueYear != "" &&
		p.expirationYear != "" &&
		p.height != "" &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportID != "") {
		return false
	}

	// validate birth year
	by, err := strconv.Atoi(p.birthYear)
	if err != nil || by < 1920 || by > 2020 {
		return false
	}

	// validate issue year
	iy, err := strconv.Atoi(p.issueYear)
	if err != nil || iy < 2010 || iy > 2020 {
		return false
	}

	// validate expiration year
	ey, err := strconv.Atoi(p.expirationYear)
	if err != nil || ey < 2020 || ey > 2030 {
		return false
	}

	// validate height
	var height int
	var unit string
	_, err = fmt.Sscanf(p.height, "%d%s", &height, &unit)
	if err != nil {
		return false
	}
	switch unit {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}

	// validate hair color
	if len(p.hairColor) != 7 {
		return false
	}
	if !strings.HasPrefix(p.hairColor, "#") {
		return false
	}
	hcHex := strings.TrimPrefix(p.hairColor, "#")
	_, err = hex.DecodeString(hcHex)
	if err != nil {
		return false
	}

	// validate eye color
	switch p.eyeColor {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	// validate passport id
	if len(p.passportID) != 9 {
		return false
	}
	_, err = strconv.Atoi(p.passportID)
	if err != nil {
		return false
	}

	return true
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
