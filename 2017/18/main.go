package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	regs   map[string]int64
	i      int64
	instrs []string
	sent   int

	rcvq []int64
	sndp *Program

	id int64
}

func New(pid int64) *Program {
	p := &Program{regs: map[string]int64{}, id: pid}
	p.Set("p", pid)

	return p
}

func (p *Program) Get(name string) int64 {
	if v, ok := p.regs[name]; ok {
		return v
	}
	return 0
}

func (p *Program) Set(name string, value int64) {
	if name == "" {
		panic("empty name")
	}
	p.regs[name] = value
}

func (p *Program) Connect(other *Program) {
	p.sndp = other
	other.sndp = p
}

func (p *Program) Send(v int64) {
	p.rcvq = append(p.rcvq, v)
}

func (p *Program) Receive() *int64 {
	if len(p.rcvq) == 0 {
		return nil
	}
	v := p.rcvq[0]
	p.rcvq = p.rcvq[1:]

	return &v
}

func (p *Program) Load(r io.Reader) {
	p.instrs = nil

	s := bufio.NewScanner(r)
	for s.Scan() {
		p.instrs = append(p.instrs, s.Text())
	}
}

func (p *Program) Step() (blocked bool) {
	instr := p.instrs[p.i]
	args := strings.Split(instr[4:], " ")
	instr = instr[:3]

	// Register args 1 and 2
	var r1, r2 string
	// Immediate args 1 and 2
	var im1, im2 int64
	// Indirect arg 1 and 2
	var in1, in2 int64
	// Args 1 and 2 values, immediate or indirect
	var v1, v2 int64

	if len(args) == 2 {
		v, err := strconv.Atoi(args[1])
		if err == nil {
			im2 = int64(v)
			v2 = im2
		} else {
			r2 = args[1]
		}
	}
	v, err := strconv.Atoi(args[0])
	if err == nil {
		im1 = int64(v)
		v1 = im1
	} else {
		r1 = args[0]
	}
	if r1 != "" {
		in1 = p.Get(r1)
		v1 = in1
	}
	if r2 != "" {
		in2 = p.Get(r2)
		v2 = in2
	}

	//fmt.Println("pid:", p.id, "full:", p.instrs[p.i], "instr:", instr, "r1:", r1, "r2:", r2, "in1:", in1, "in2:", in2, "im2:", im2, "v2:", v2)
	//fmt.Println("BEFORE regs:", p.regs, "rcvq:", p.rcvq, "other rcvq:", p.sndp.rcvq)

	switch instr {
	case "snd":
		p.sndp.Send(v1)
		p.sent++
		//fmt.Println("pid:", p.id, "sent; total", p.sent)
	case "set":
		p.Set(r1, v2)
	case "add":
		p.Set(r1, in1+v2)
	case "mul":
		p.Set(r1, in1*v2)
	case "mod":
		p.Set(r1, in1%v2)
	case "rcv":
		v := p.Receive()
		if v == nil {
			// we're blocked
			//fmt.Println("pid:", p.id, "is blocked!")
			return true
		}
		p.Set(r1, *v)
	case "jgz":
	default:
		panic(instr)
	}

	//fmt.Println("AFTER  regs:", p.regs, "rcvq:", p.rcvq, "other rcvq:", p.sndp.rcvq)

	if instr == "jgz" && in1 > 0 {
		p.i += v2
	} else {
		p.i++
	}

	return false
}

func RunUntilDeadlocked(p0, p1 *Program) int {
	for !p0.Step() || !p1.Step() {
	}
	//for {
	//	p0Blocked, p1Blocked := p0.Step(), p1.Step()
	//	if p0Blocked && p1Blocked {
	//		return p1.sent
	//	}
	//}
	return p1.sent
}

func main() {
	p0 := New(0)
	p1 := New(1)

	var buf bytes.Buffer
	t := io.TeeReader(os.Stdin, &buf)
	p0.Load(t)
	p1.Load(&buf)

	p0.Connect(p1)

	//fmt.Println("p0", p0.instrs)
	//fmt.Println("p1", p1.instrs)

	fmt.Println("p1 sent:", RunUntilDeadlocked(p0, p1))
}
