package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addOne(v int) int {
	return v + 1
}

func subThreeOrMoreElseAddOne(v int) int {
	if v >= 3 {
		return v - 1
	}
	return v + 1
}

func jmpsToExit(jmps []int, mod func(int) int) int {
	tot := 0
	pc := 0
	for {
		if pc < 0 || pc >= len(jmps) {
			return tot
		}
		jmp := jmps[pc]
		jmps[pc] = mod(jmps[pc])
		pc += jmp
		tot++
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var jmps []int
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		jmps = append(jmps, v)
	}
	// Copy since we modify it in place.
	jmps1 := append([]int(nil), jmps...)

	fmt.Println("jmpsToExit(addOne):", jmpsToExit(jmps, addOne))
	fmt.Println("jmpsToExit(subThreeOrMoreElseAddOne):",
		jmpsToExit(jmps1, subThreeOrMoreElseAddOne))
}
