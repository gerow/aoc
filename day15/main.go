package main

import "fmt"

const (
	NormalIterations   = 40000000
	CriteriaIterations = 5000000
)

type Generator struct {
	factor   int
	next     int
	criteria int
}

func New(factor, next, criteria int) *Generator {
	return &Generator{factor, next, criteria}
}

func NewA(next int) *Generator {
	return New(16807, next, 0)
}

func NewB(next int) *Generator {
	return New(48271, next, 0)
}

func NewAWithCriteria(next int) *Generator {
	return New(16807, next, 4)
}

func NewBWithCriteria(next int) *Generator {
	return New(48271, next, 8)
}

func (g *Generator) Next() int {
again:
	n := g.next
	g.next = n * g.factor % 2147483647
	if g.criteria != 0 && n%g.criteria != 0 {
		goto again
	}
	return n
}

func Match(a, b *Generator, trials int) int {
	var matches int
	for n := 0; n < trials; n++ {
		av, bv := a.Next(), b.Next()
		if av&0xffff == bv&0xffff {
			matches++
		}
	}
	return matches
}

func main() {
	a := NewA(516)
	b := NewB(190)
	fmt.Println("Matches:", Match(a, b, NormalIterations))

	ac := NewAWithCriteria(516)
	bc := NewBWithCriteria(190)
	fmt.Println("Matches with criteria:", Match(ac, bc, CriteriaIterations))
}
