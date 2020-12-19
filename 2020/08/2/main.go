package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type cpu struct {
	acc, pc      int
	instructions []instruction
}

type instruction struct {
	op  op
	arg int
}

type op int

const (
	acc op = iota
	jmp
	nop
)

func (c *cpu) clone() *cpu {
	var c2 cpu = *c
	c2.instructions = make([]instruction, len(c.instructions))
	copy(c2.instructions, c.instructions)

	return &c2
}

func (c *cpu) load(r io.Reader) error {
	s := bufio.NewScanner(os.Stdin)
	var line int
	for s.Scan() {
		line++
		var ins instruction
		var opStr string
		_, err := fmt.Sscanf(s.Text(), "%s %d", &opStr, &ins.arg)
		if err != nil {
			return fmt.Errorf("malformed instruction %q on line %d: %w", s.Text(), line, err)
		}
		switch opStr {
		case "acc":
			ins.op = acc
		case "jmp":
			ins.op = jmp
		case "nop":
			ins.op = nop
		default:
			return fmt.Errorf("unknown operation %q in instruction %q on line %q", opStr, s.Text(), line)
		}
		c.instructions = append(c.instructions, ins)
	}
	return s.Err()
}

func (c *cpu) step() error {
	if c.pc > len(c.instructions) {
		return fmt.Errorf("pc %d walks off the end of our instructions (len %d)", c.pc, len(c.instructions))
	}
	ins := c.instructions[c.pc]

	switch ins.op {
	case acc:
		c.acc += ins.arg
		c.pc++
	case jmp:
		c.pc += ins.arg
	case nop:
		c.pc++
	default:
		panic(fmt.Sprintf("unknown operation %d", ins.op))
	}

	return nil
}

var loopErr = errors.New("instruction seen multiple times, execution will never terminate")

func (c *cpu) stepUntilTermates() error {
	seen := make(map[int]bool)
	for {
		if seen[c.pc] {
			return loopErr
		}
		// normal termination
		if c.pc == len(c.instructions) {
			return nil
		}

		seen[c.pc] = true
		if err := c.step(); err != nil {
			return err
		}
	}
}

func main() {
	orig := &cpu{}
	if err := orig.load(os.Stdin); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(orig.instructions); i++ {
		if orig.instructions[i].op != jmp && orig.instructions[i].op != nop {
			continue
		}
		cpu := orig.clone()
		var ins *instruction = &cpu.instructions[i]
		switch ins.op {
		case jmp:
			ins.op = nop
		case nop:
			ins.op = jmp
		default:
			panic("ops should be exhaustive")
		}

		if err := cpu.stepUntilTermates(); err != nil {
			if err == loopErr {
				continue
			}
			log.Fatal(err)
		}

		fmt.Println(cpu.acc)
		os.Exit(0)
	}

	log.Fatal("failed to find mutation that causes termination")
}
