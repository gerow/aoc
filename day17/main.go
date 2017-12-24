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

func main() {
	fmt.Println("Value after last:", AfterLast(355))
}
