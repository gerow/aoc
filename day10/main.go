package main

import "fmt"

type Knot struct {
	String [256]int
	Pos    int
	Skip   int
}

func New() *Knot {
	var k Knot
	for i := 0; i < len(k.String); i++ {
		k.String[i] = i
	}
	return &k
}

func (k *Knot) Apply(length int) {
	if length > 1 {
		i, j := k.Pos, (k.Pos+length-1)%len(k.String)
		for {
			k.String[i], k.String[j] = k.String[j], k.String[i]
			i++
			i %= len(k.String)
			if i == j {
				break
			}
			j--
			if j < 0 {
				j = len(k.String) - 1
			}
			if i == j {
				break
			}
		}
	}

	k.Pos += length + k.Skip
	k.Pos %= len(k.String)
	k.Skip++
}

func main() {
	lens := []int{212, 254, 178, 237, 2, 0, 1, 54, 167, 92, 117, 125, 255, 61, 159, 164}
	k := New()
	//k.Pos = 255
	fmt.Println("Before:", k)
	for _, l := range lens {
		fmt.Println("Applying", l)
		k.Apply(l)
		fmt.Println(k)
	}
	fmt.Println("After:", k)
	fmt.Println("String[0]*String[1]:", k.String[0]*k.String[1])
}
