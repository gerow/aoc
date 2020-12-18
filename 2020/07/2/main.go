package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type containsRule struct {
	number int
	color  string
}

type rule struct {
	color    string
	contains []containsRule
}

func parseContainsRule(s string) (containsRule, error) {
	rest := strings.TrimSpace(strings.TrimSuffix(s, "."))
	rest = strings.TrimSpace(strings.TrimSuffix(rest, "bags"))
	rest = strings.TrimSpace(strings.TrimSuffix(rest, "bag"))
	sp := strings.SplitN(rest, " ", 2)
	if len(sp) != 2 {
		return containsRule{}, fmt.Errorf("containsRule %q is malformed", s)
	}
	var cr containsRule
	var err error
	cr.number, err = strconv.Atoi(sp[0])
	if err != nil {
		return containsRule{}, fmt.Errorf("containsRule %q is malformed: %w", s, err)
	}
	cr.color = sp[1]

	return cr, nil
}

func parseRule(s string) (*rule, error) {
	var r rule
	sl := strings.SplitN(s, "bags contain", 2)
	if len(sl) != 2 {
		return nil, fmt.Errorf("malformed rule %q missing \"bags contain\" directive", s)
	}
	r.color = strings.TrimSpace(sl[0])
	rest := strings.TrimSpace(sl[1])
	if rest == "no other bags." {
		return &r, nil
	}
	for _, crs := range strings.Split(rest, ",") {
		cr, err := parseContainsRule(strings.TrimSpace(crs))
		if err != nil {
			return nil, fmt.Errorf("malformed rule %q: %w", s, err)
		}
		r.contains = append(r.contains, cr)
	}

	return &r, nil
}

func bagsIn(c string, rules map[string]*rule) int {
	var n int
	for _, r := range rules[c].contains {
		n += r.number
		n += bagsIn(r.color, rules) * r.number
	}

	return n
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	rules := make(map[string]*rule)
	for s.Scan() {
		r, err := parseRule(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		rules[r.color] = r
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(bagsIn("shiny gold", rules))
}
