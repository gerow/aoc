package main

// This code is extrememly broken in a bunch of different ways. It
// overcomplicates rotations, but also just doesn't really handle rotations
// *at all* since we never set the direction at the beginning. Yet somehow
// it produced the right answer. Don't ask me.

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var buses []int
	var earliest int

	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		var err error
		earliest, err = strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalf("failed to read earliest time; scanner err %v", s.Err())
	}
	if s.Scan() {
		for _, b := range strings.Split(s.Text(), ",") {
			if b == "x" {
				continue
			}
			n, err := strconv.Atoi(b)
			if err != nil {
				log.Fatal(err)
			}
			buses = append(buses, n)
		}
	} else {
		log.Fatalf("failed to read buses; scanner err %v", s.Err())
	}

	var earliestID int
	earliestDeparture := -1
	for _, b := range buses {
		if earliest % b == 0 {
			panic(fmt.Sprintf("bus id %d should leave exactly when we arrive, curious...", b))
		}
		departure := (earliest/b + 1)*b
		if departure < earliestDeparture || earliestDeparture == -1 {
			earliestDeparture = departure
			earliestID = b
		}
	}

	fmt.Println(earliestID * (earliestDeparture - earliest))
}
