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
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to convert %s to an int: %v", s.Text(), err)
		}
		for _, no := range nums {
			if n+no == 2020 {
				fmt.Println(n * no)
				os.Exit(0)
			}
		}
		nums = append(nums, n)
	}

	log.Fatalf("failed to find pair that sums to 2020")
}
