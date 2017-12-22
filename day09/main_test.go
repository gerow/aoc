package main

import (
	"strings"
	"testing"
)

func TestReference(t *testing.T) {
	for _, st := range []struct {
		s    string
		want int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	} {
		r := strings.NewReader(st.s)
		if got := Score(r); got != st.want {
			t.Errorf("wrong Score for %q; want %d got %d", st.s, st.want, got)
		}
	}
}
