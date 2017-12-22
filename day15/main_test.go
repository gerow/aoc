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
	} {
		for i, expected := range gt.expected {
			if actual := gt.g.Next(); actual != expected {
				t.Errorf("wrong Next at index %d; want %v got %v", i, expected, actual)
			}
		}
	}

	a := NewA(1092455)
	b := NewB(430625591)
	if v := Match(a, b, 40000000); v != 588 {
		t.Errorf("wrong match; want 588 got %v", v)
	}
}
