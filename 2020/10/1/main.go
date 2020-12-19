package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var adapters []int

	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to read %q as a number: %v", s.Text(), err)
		}
		adapters = append(adapters, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// As far as I can tell determining the correct order is simply a matter of sorting the adapters.
	sort.Ints(adapters)
	var j1, j3 int
	var prev int
	for _, v := range adapters {
		diff := v - prev
		switch diff {
		case 1:
			j1++
		case 2:
		case 3:
			j3++
		default:
			log.Fatalf("joltage difference of %d required, only 1, 2, or 3 should be acceptable", diff)
		}
		prev = v
	}
	// add a final jump to j3 for our device
	j3++

	fmt.Println(j1 * j3)
}
