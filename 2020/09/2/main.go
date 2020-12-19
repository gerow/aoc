package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type xmas struct {
	prev []int
}

func (x *xmas) valid(n int) bool {
	if len(x.prev) != 25 {
		panic(fmt.Sprintf("somehow x.prev ended up being length %d", len(x.prev)))
	}

	// I think we can do this way *way* faster (but try it slow for now).
	for i := 0; i < len(x.prev)-1; i++ {
		for j := i + 1; j < len(x.prev); j++ {
			if x.prev[i]+x.prev[j] == n {
				return true
			}
		}
	}

	return false
}

func (x *xmas) extend(n int) {
	x.prev = append(x.prev[1:], n)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var x xmas
	var nums []int
	for i := 0; i < 25; i++ {
		if !s.Scan() {
			if s.Err() != nil {
				log.Fatal(s.Err())
			}
			log.Fatalf("failed to read full preamble; read %d of 25 numbers", i)
		}
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to read %q as a number on line %d: %v", s.Text(), i+1, err)
		}
		nums = append(nums, n)
		x.prev = append(x.prev, n)
	}

	var invalid int
	var found bool
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to read %q as a number: %v", s.Text(), err)
		}
		nums = append(nums, n)
		if !x.valid(n) {
			found = true
			invalid = n
			break
		}
		x.extend(n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	if !found {
		log.Fatal("failed to find an invalid number")
	}

	// slurp down the rest of the numbers (streaming stuff in complicates stuff for part 2, but I didn't want to rewrite everything...)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("failed to read %q as a number: %v", s.Text(), err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	var beginning, end int
	for i := 0; i < len(nums) - 1; i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == invalid {
				beginning = i
				end = j
				goto found
			}
			if sum > invalid {
				break
			}
		}
	}
	log.Fatalf("failed to find run that summs to invalid number %d", invalid)

found:
	smallest, largest := nums[beginning], nums[beginning]

	for i := beginning + 1; i < end + 1; i++ {
		if nums[i] < smallest {
			smallest = nums[i]
		}
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	fmt.Println(smallest + largest)
}
