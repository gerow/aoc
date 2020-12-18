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

const lookingFor = "shiny gold"

func main() {
	s := bufio.NewScanner(os.Stdin)
	var rules []*rule
	for s.Scan() {
		r, err := parseRule(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		rules = append(rules, r)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// create an index that maps us from the containing bag's color to the outer bag's color
	isContainedIn := map[string][]string{}
	for _, r := range rules {
		for _, cr := range r.contains {
			e := isContainedIn[cr.color]
			e = append(e, r.color)
			isContainedIn[cr.color] = e
		}
	}

	startingPoints := make(map[string]bool)
	added := map[string]bool{lookingFor: true}
	toCheck := []string{lookingFor}

	for len(toCheck) != 0 {
		c := toCheck[0]
		toCheck = toCheck[1:]

		for _, ici := range isContainedIn[c] {
			startingPoints[ici] = true
			if !added[ici] {
				toCheck = append(toCheck, ici)
				added[ici] = true
			}
		}
	}

	fmt.Println(len(startingPoints))
}
