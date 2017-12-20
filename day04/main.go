package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func identity(s string) string {
	return s
}

// Finds the "canonical" anagram. That is: the anagram sorted by rune value.
func canonicalAnagram(s string) string {
	var ss []string
	for _, r := range s {
		ss = append(ss, string(r))
	}
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func numValid(r io.Reader, munger func(string) string) int {
	s := bufio.NewScanner(r)
	numValid := 0
	for s.Scan() {
		words := strings.Split(s.Text(), " ")
		used := map[string]bool{}
		valid := true
		for _, w := range words {
			w = munger(w)
			if _, ok := used[w]; ok {
				valid = false
				break
			}
			used[w] = true
		}
		if valid {
			numValid++
		}
	}

	return numValid
}

func main() {
	var buf bytes.Buffer
	t := io.TeeReader(os.Stdin, &buf)

	go fmt.Println("numValid:", numValid(t, identity))
	fmt.Println("numValidAnagram:", numValid(&buf, canonicalAnagram))
}
