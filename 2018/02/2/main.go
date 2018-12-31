package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// diff returns a slice of ints indicating the indicies where a and b differ
func diff(a, b string) []int {
	if len(a) != len(b) {
		panic("diff args are of different length")
	}

	var d []int
	for i := range a {
		if a[i] != b[i] {
			d = append(d, i)
		}
	}

	return d
}

func main() {
	var ids []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ids = append(ids, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	for i, a := range ids {
		for _, b := range ids[i+1:] {
			d := diff(a, b)
			if len(d) != 1 {
				continue
			}
			idx := d[0]
			fmt.Println(a[:idx] + a[idx+1:])
			os.Exit(0)
		}
	}

	log.Fatal("failed to find ids that diff by one letter")
}
