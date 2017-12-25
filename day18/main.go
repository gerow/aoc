package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type registers struct {
	regs map[string]int
}

func (r *registers) Get(name string) int {
	if v, ok := r.regs[name]; ok {
		return v
	}
	return 0
}

func (r *registers) Set(name string, value int) {
	r.regs[name] = value
}

type Machine struct {
	r      registers
	i      int
	freq   int
	instrs []string
}

func New() *Machine {
	return &Machine{r: registers{regs: map[string]int{}}}
}

func (m *Machine) Load(r io.Reader) {
	m.instrs = nil

	s := bufio.NewScanner(r)
	for s.Scan() {
		m.instrs = append(m.instrs, s.Text())
	}
}

func (m *Machine) Recover() int {
	for {
		instr := m.instrs[m.i]
		args := strings.Split(instr[4:], " ")
		instr = instr[:3]

		// Register args 1 and 2
		var r1, r2 string
		// Immediate arg 2
		var im2 int
		// Indirect arg 1 and 2
		var in1, in2 int
		// Arg 2 value, immediate or indirect
		var v2 int

		r1 = args[0]
		if len(args) == 2 {
			v, err := strconv.Atoi(args[1])
			if err == nil {
				im2 = v
				v2 = im2
			} else {
				r2 = args[1]
			}
		}
		if r1 != "" {
			in1 = m.r.Get(r1)
		}
		if r2 != "" {
			in2 = m.r.Get(r2)
			v2 = in2
		}

		switch instr {
		case "snd":
			m.freq = in1
		case "set":
			m.r.Set(r1, v2)
		case "add":
			m.r.Set(r1, in1+v2)
		case "mul":
			m.r.Set(r1, in1*v2)
		case "mod":
			m.r.Set(r1, in1%v2)
		case "rcv":
			if in1 != 0 {
				return m.freq
			}
		}

		//fmt.Println("full:", m.instrs[m.i], "instr:", instr, "r1:", r1, "r2:", r2, "in1:", in1, "in2:", in2, "im2:", im2, "v2:", v2)
		//fmt.Println("regs:", m.r)

		if instr == "jgz" && in1 > 0 {
			m.i += v2
		} else {
			m.i++
		}
	}
}

func main() {
	m := New()
	m.Load(os.Stdin)

	fmt.Println("Recover:", m.Recover())
}
