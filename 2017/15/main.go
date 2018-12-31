package main

import "fmt"

const (
	NormalIterations   = 40000000
	CriteriaIterations = 5000000
)

type Generator struct {
	factor   int
	last     int
	criteria int
}

func New(factor, last, criteria int) *Generator {
	return &Generator{factor, last, criteria}
}

func NewA(last int) *Generator {
	return New(16807, last, 0)
}

func NewB(last int) *Generator {
	return New(48271, last, 0)
}

func NewAWithCriteria(last int) *Generator {
	return New(16807, last, 4)
}

func NewBWithCriteria(last int) *Generator {
	return New(48271, last, 8)
}

func (g *Generator) Next() int {
again:
	g.last = g.last * g.factor % 2147483647
	if g.criteria != 0 && g.last%g.criteria != 0 {
		goto again
	}
	return g.last
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
