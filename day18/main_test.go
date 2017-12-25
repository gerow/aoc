package main

import (
	"strings"
	"testing"
)

func TestReference(t *testing.T) {
	for _, rt := range []struct {
		instrs string
		want   int
	}{
		{`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
`, 4},
	} {
		m := New()
		m.Load(strings.NewReader(rt.instrs))
		if got := m.Recover(); got != rt.want {
			t.Errorf("wrong Recover for %q: got %d want %d", rt.instrs, got, rt.want)
		}
	}
}
