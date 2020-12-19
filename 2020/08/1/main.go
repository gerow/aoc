package main

import (
	"bufio"
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

func (c *cpu) stepUntilRepeat() error {
	seen := make(map[int]bool)
	for !seen[c.pc] {
		seen[c.pc] = true
		if err := c.step(); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var cpu cpu
	if err := cpu.load(os.Stdin); err != nil {
		log.Fatal(err)
	}
	if err := cpu.stepUntilRepeat(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cpu.acc)
}
