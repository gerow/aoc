package main

import "testing"

func TestReference(t *testing.T) {
	for _, gt := range []struct {
		g        *Generator
		expected []int
	}{
		{NewA(1092455), []int{
			1092455,
			1181022009,
			245556042,
			1744312007,
			1352636452,
		}},
		{NewB(430625591), []int{
			430625591,
			1233683848,
			1431495498,
			137874439,
			285222916,
		}},
		{NewAWithCriteria(1092455), []int{
			1352636452,
			1992081072,
			530830436,
			1980017072,
			740335192,
		}},
		{NewBWithCriteria(430625591), []int{
			1233683848,
			862516352,
			1159784568,
			1616057672,
			412269392,
		}},
	} {
		for i, expected := range gt.expected {
			if actual := gt.g.Next(); actual != expected {
				t.Errorf("wrong Next at index %d; want %v got %v", i, expected, actual)
			}
		}
	}

	a := NewA(1092455)
	b := NewB(430625591)
	if v := Match(a, b, NormalIterations); v != 588 {
		t.Errorf("wrong match; want 588 got %v", v)
	}

	ac := NewAWithCriteria(1092455)
	bc := NewBWithCriteria(430625591)
	if v := Match(ac, bc, CriteriaIterations); v != 309 {
		t.Errorf("wrong match; want 309 got %v", v)
	}
}
