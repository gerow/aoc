package main

import (
	"strings"
	"testing"
)

func TestReference(t *testing.T) {
	for _, st := range []struct {
		s            string
		score        int
		garbageCount int
	}{
		{"{}", 1, 0},
		{"{{{}}}", 6, 0},
		{"{{},{}}", 5, 0},
		{"{{{},{},{{}}}", 16, 0},
		{"{<a>,<a>,<a>,<a>}", 1, 4},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
	} {
		r := strings.NewReader(st.s)
		if got := Score(r); got != st.score {
			t.Errorf("wrong Score for %q; want %d got %d", st.s, st.score, got)
		}
		r.Reset(st.s)
		if got := GarbageCount(r); got != st.garbageCount {
			t.Errorf("wrong GarbageCount for %q: want %d got %d", st.s, st.garbageCount, got)
		}
	}
}
