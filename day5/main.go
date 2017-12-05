package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func jmpsToExit(jmps []int) int {
	tot := 0
	pc := 0
	for {
		if pc < 0 || pc >= len(jmps) {
			return tot
		}
		jmp := jmps[pc]
		jmps[pc]++
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

	fmt.Println("jmpsToExit:", jmpsToExit(jmps))
}
