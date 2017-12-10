package main

import (
	"bufio"
	"fmt"
	"os"
)

type Registers struct {
	R map[string]int
}

func (r *Registers) Read(reg string) int {
	if v, ok := r.R[reg]; ok {
		return v
	}
	return 0
}

func (r *Registers) Inc(tok, reg string, val int) {
	if tok != "inc" && tok != "dec" {
		panic(tok)
	}
	if tok == "dec" {
		val = -val
	}

	if _, ok := r.R[reg]; ok {
		r.R[reg] += val
	} else {
		r.R[reg] = val
	}
}

func (r *Registers) Condition(reg, cond string, val int) bool {
	switch cond {
	case "<":
		return r.Read(reg) < val
	case ">":
		return r.Read(reg) > val
	case "<=":
		return r.Read(reg) <= val
	case ">=":
		return r.Read(reg) >= val
	case "==":
		return r.Read(reg) == val
	case "!=":
		return r.Read(reg) != val
	}
	panic(cond)
}

func (r *Registers) Largest() int {
	var largest int
	for _, v := range r.R {
		largest = v
		break
	}
	for _, v := range r.R {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	r := Registers{map[string]int{}}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// wReg incTok val if rReg cmpTok cmpVal
		var wReg, incTok, rReg, cmpTok string
		var val, cmpVal int
		if _, err := fmt.Sscanf(s.Text(), "%s %s %d if %s %s %d", &wReg, &incTok, &val, &rReg, &cmpTok, &cmpVal); err != nil {
			panic(err)
		}
		if r.Condition(rReg, cmpTok, cmpVal) {
			r.Inc(incTok, wReg, val)
		}
	}
	fmt.Println("largest:", r.Largest())
}
