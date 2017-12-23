package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanCommas(t *testing.T) {
	for _, st := range []struct {
		s    string
		want []string
	}{
		{"a,b,c,d,e", []string{"a", "b", "c", "d", "e"}},
		{"ab,cd,ef,gh,ij", []string{"ab", "cd", "ef", "gh", "i"}},
	} {
		r := strings.NewReader(st.s)
		s := bufio.NewScanner(r)
		s.Split(ScanCommas)
		for i := 0; s.Scan(); i++ {
			if got := s.Text(); got != st.want[i] {
				t.Errorf("wrong value for %q at index %d; want %s got %s", st.s, i, st.want[i], got)
			}
		}
	}
}
