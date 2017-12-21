package main

import (
	"fmt"
	"math/bits"

	"github.com/gerow/aoc/day10/knot"
)

const totalRows = 128

func main() {
	in := "hxtvlmkl"

	used := 0
	for r := 0; r < totalRows; r++ {
		h := knot.Sum([]byte(fmt.Sprintf("%s-%d", in, r)))
		for _, b := range h {
			used += bits.OnesCount8(b)
		}
	}
	fmt.Println("Used:", used)
}
