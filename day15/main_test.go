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

	for _, mt := range []struct {
		sa, sb, match, criteria int
	}{
		{1092455, 430625591, 588, 309},
		{516, 190, 597, 303},
	} {
		a := NewA(mt.sa)
		b := NewB(mt.sb)
		if v := Match(a, b, NormalIterations); v != mt.match {
			t.Errorf("wrong match for (%d, %d); want %d got %d", mt.sa, mt.sb, mt.match, v)
		}

		ac := NewAWithCriteria(mt.sa)
		bc := NewBWithCriteria(mt.sb)
		if v := Match(ac, bc, CriteriaIterations); v != mt.criteria {
			t.Errorf("wrong criteria match for (%d, %d); want %d got %d", mt.sa, mt.sb, mt.criteria, v)
		}
	}
}
