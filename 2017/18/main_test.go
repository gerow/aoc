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
		{`snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`, 3},
	} {
		p0 := New(0)
		p0.Load(strings.NewReader(rt.instrs))

		p1 := New(1)
		p1.Load(strings.NewReader(rt.instrs))

		p0.Connect(p1)
		if got := RunUntilDeadlocked(p0, p1); got != rt.want {
			t.Errorf("wrong RunUntilDeadlock for %q: got %d want %d", rt.instrs, got, rt.want)
		}
	}
}
