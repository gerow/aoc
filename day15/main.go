package main

import "fmt"

type Generator struct {
	factor int
	next   int
}

func New(factor, next int) *Generator {
	return &Generator{factor, next}
}

func NewA(next int) *Generator {
	return New(16807, next)
}

func NewB(next int) *Generator {
	return New(48271, next)
}

func (g *Generator) Next() int {
	n := g.next
	g.next = n * g.factor % 2147483647
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
	fmt.Println("Matches:", Match(a, b, 40000000))
}
