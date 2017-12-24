package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScanCommas(t *testing.T) {
	for _, st := range []struct {
		s    string
		want []string
	}{
		{"s1", []string{"s1"}},
		{"a,b,c,d,e", []string{"a", "b", "c", "d", "e"}},
		{"ab,cd,ef,gh,ij", []string{"ab", "cd", "ef", "gh", "ij"}},
		{"ab,cd,ef,gh,ij\n", []string{"ab", "cd", "ef", "gh", "ij"}},
	} {
		r := strings.NewReader(st.s)
		s := bufio.NewScanner(r)
		s.Split(ScanCommas)
		for i, want := range st.want {
			s.Scan()
			if got := s.Text(); got != want {
				t.Errorf("wrong value for %q at index %d; want %q got %q", st.s, i, want, got)
			}
		}
		if s.Err() != nil {
			fmt.Errorf("unexpected scan error: %v", s.Err())
		}
	}
}

func TestReference(t *testing.T) {
	for _, rt := range []struct {
		moves, want string
	}{
		{"s1", "eabcd"},
		{"s2", "deabc"},
		{"s1,x3/4,pe/b", "baedc"},
	} {
		s := New(len(rt.want))
		r := strings.NewReader(rt.moves)
		moves := ParseMoves(r)
		s.ApplyMoves(moves)
		if got := string(s.Programs); got != rt.want {
			t.Errorf("wrong result for moves %q; want %q got %q", rt.moves, rt.want, got)
		}
	}
}
