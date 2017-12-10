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
	ChildrenNames []string
	Children      []*Program
}

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

	for _, p := range programs {
		fmt.Printf("%+v\n", *p)
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
}
