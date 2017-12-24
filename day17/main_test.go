package main

import (
	"testing"
)

func TestReference(t *testing.T) {
	for _, rt := range []struct {
		step, want int
	}{
		{3, 638},
	} {
		if got := AfterLast(rt.step); got != rt.want {
			t.Errorf("wrong AfterLast for %d: got %d want %d", rt.step, got, rt.want)
		}
	}
}
