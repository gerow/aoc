package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Banks [16]int

func Rebalance(b Banks) int {
	m := map[Banks]bool{}
	n := 0
	for {
		if _, ok := m[b]; ok {
			return n
		}
		m[b] = true
		n++

		// Find biggest
		iBig := 0
		vBig := b[iBig]
		for i := 1; i < len(b); i++ {
			if b[i] > vBig {
				iBig = i
				vBig = b[iBig]
			}
		}
		// Redistribute
		bucket := b[iBig]
		b[iBig] = 0
		i := (iBig + 1) % len(b)
		for bucket > 0 {
			b[i]++
			bucket--
			i = (i + 1) % len(b)
		}
	}
}

func main() {
	s := "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11"
	var b Banks
	for i, e := range strings.Split(s, "\t") {
		ei, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		b[i] = ei
	}

	fmt.Println(b)
	fmt.Println("iterations:", Rebalance(b))
}
