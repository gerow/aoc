package main

import (
	"container/ring"
	"fmt"
)

func AfterLast(step int) int {
	r := ring.New(1)
	r.Value = 0

	for n := 1; n < 2018; n++ {
		for s := 0; s < step; s++ {
			r = r.Next()
		}
		r.Link(&ring.Ring{Value: n})
		r = r.Next()
	}
	return r.Next().Value.(int)
}

func AfterZero(step int) int {
	var i, afterZero int
	for n := 1; n < 50000001; n++ {
		i = (i + step) % n
		if i == 0 {
			afterZero = n
		}
		i++
	}
	return afterZero
}

func main() {
	fmt.Println("Value after last:", AfterLast(355))
	fmt.Println("Value after zero:", AfterZero(355))
}
