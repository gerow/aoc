package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Program struct {
	Name          string
	Weight        int
	SumWeight     int
	ChildrenNames []string
	Children      []*Program
}

func PopulateSumWeight(p *Program) {
	for _, c := range p.Children {
		PopulateSumWeight(c)
	}
	p.SumWeight = p.Weight
	for _, c := range p.Children {
		p.SumWeight += c.SumWeight
	}
}

func continuePrint(p *Program, depth int) {
	fmt.Print(strings.Repeat("  ", depth))
	fmt.Printf("%s %d: ", p.Name, p.SumWeight)
	for _, c := range p.Children {
		fmt.Printf("| %s %d", c.Name, c.SumWeight)
	}
	fmt.Printf("\n")
	for _, c := range p.Children {
		continuePrint(c, depth+1)
	}
}

func Print(p *Program) {
	continuePrint(p, 0)
}

/*
func FindBad(p *Program) (*Program, int) {
	if len(p.Children) == 0 {
		return nil, 0
	}
	expected := p.Children[0].SumWeight
	allGood := true
	for _, c := range p.Children {
		if c.SumWeight != expected {
			allGood = false
			break
		}
	}
	if allGood {
		return nil, 0
	}

	for _, c := range p.Children {
		p, n := FindBad(c)
		if p != nil {
			return p, n
		}
	}
	panic("this should never happen")
}
*/

func main() {
	s := bufio.NewScanner(os.Stdin)
	programs := []*Program{}
	for s.Scan() {
		parts := strings.Split(s.Text(), " -> ")
		if len(parts) != 1 && len(parts) != 2 {
			panic(s.Text())
		}
		hdr := parts[0]
		var children string
		if len(parts) == 2 {
			children = parts[1]
		}
		var name string
		var weight int
		_, err := fmt.Sscanf(hdr, "%s (%d)", &name, &weight)
		if err != nil {
			panic(err)
		}
		p := &Program{
			Name:   name,
			Weight: weight,
		}
		if children != "" {
			p.ChildrenNames = strings.Split(children, ", ")
		}
		programs = append(programs, p)
	}

	// Now for each program look for its parent. This is O(n^2) but could be improved.
	var root *Program

	for _, p := range programs {
		foundParent := false
		for _, candidate := range programs {
			for _, childName := range candidate.ChildrenNames {
				if p.Name == childName {
					foundParent = true
					candidate.Children = append(candidate.Children, p)
					break
				}
			}
			if foundParent {
				break
			}
		}
		if !foundParent {
			if root != nil {
				panic("found multiple roots???")
			}
			root = p
		}
	}
	fmt.Printf("root: %+v\n", *root)
	PopulateSumWeight(root)
	// The answer is 1226, unforunately I couldn't devise a good way to do
	// that programatically, I just visually scanned through the graph.
	Print(root)
}
