package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PID int

type Program struct {
	Pipes []PID
}

func ConnectedSize(programs map[PID]Program, pid PID) int {
	return connectedSize(programs, pid, map[PID]bool{})
}

func connectedSize(programs map[PID]Program, pid PID, seen map[PID]bool) int {
	sum := 1 // self
	seen[pid] = true
	for _, ppid := range programs[pid].Pipes {
		// We've seen it already so skip
		if _, ok := seen[ppid]; ok {
			continue
		}
		sum += connectedSize(programs, ppid, seen)
	}
	return sum
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	programs := map[PID]Program{}
	for s.Scan() {
		sp := strings.Split(s.Text(), " <-> ")
		if len(sp) != 2 {
			panic("bad format")
		}
		pid, err := strconv.Atoi(sp[0])
		if err != nil {
			panic(err)
		}

		var pipes []PID
		for _, ps := range strings.Split(sp[1], ", ") {
			pi, err := strconv.Atoi(ps)
			if err != nil {
				panic(err)
			}
			pipes = append(pipes, PID(pi))
		}
		programs[PID(pid)] = Program{Pipes: pipes}
	}

	fmt.Println("ConnectedSize(0):", ConnectedSize(programs, 0))
}
