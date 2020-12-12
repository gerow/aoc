package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var nums []int

	for s.Scan() {
		n1, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to convert %s to an int: %v", s.Text(), err)
		}
		for i, n2 := range nums {
			for j, n3 := range nums {
				// we can't use a number twice
				if i == j {
					continue
				}
				if n1 + n2 + n3 == 2020 {
					fmt.Println(n1 * n2 * n3)
					os.Exit(0)
				}
			}
		}
		nums = append(nums, n1)
	}

	log.Fatalf("failed to find triplet that sums to 2020")
}
