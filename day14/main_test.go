package main

import "testing"

func TestReference(t *testing.T) {
	g := New("flqrgnkx")
	if u := g.Used(); u != 8108 {
		t.Errorf("wrong value for used; want 8108 got %d", u)
	}
	if g := g.Groups(); g != 1242 {
		t.Errorf("wrong value for groups; want 1242 got %d", g)
	}
	for _, ft := range []struct {
		row      int
		col      int
		expected bool
	}{
		{0, 0, true},
		{0, 1, true},
		{0, 2, false},
		{1, 0, false},
		{1, 1, true},
	} {
		if v := g.Get(ft.row, ft.col); v != ft.expected {
			t.Errorf("wrong Get at %d, %d; want %v got %v", ft.row, ft.col, ft.expected, v)
		}
	}
}
