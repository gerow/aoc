package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	numValid := 0
	for s.Scan() {
		words := strings.Split(s.Text(), " ")
		used := map[string]bool{}
		valid := true
		for _, w := range words {
			if _, ok := used[w]; ok {
				valid = false
				break
			}
			used[w] = true
		}
		if valid {
			numValid++
		}
	}
	fmt.Println("numValid:", numValid)
}
