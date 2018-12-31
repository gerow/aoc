package main

import (
	"encoding/hex"
	"fmt"

	"github.com/gerow/aoc/day10/knot"
)

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

func (k *Knot) Dense() (o [16]byte) {
	for b := 0; b < 16; b++ {
		o[b] = byte(k.String[b*16])
		for i := 1; i < 16; i++ {
			o[b] ^= byte(k.String[b*16+i])
		}
	}
	return o
}

func (k *Knot) Hex() string {
	d := k.Dense()
	return hex.EncodeToString(d[:])
}

func Lens(in string) []int {
	var o []int
	for _, b := range []byte(in) {
		o = append(o, int(b))
	}
	return append(o, []int{17, 31, 73, 47, 23}...)
}

func main() {
	lens := []int{212, 254, 178, 237, 2, 0, 1, 54, 167, 92, 117, 125, 255, 61, 159, 164}
	k := New()
	for _, l := range lens {
		k.Apply(l)
	}
	fmt.Println("String[0]*String[1]:", k.String[0]*k.String[1])
	sum := knot.Sum([]byte("212,254,178,237,2,0,1,54,167,92,117,125,255,61,159,164"))
	fmt.Println(hex.EncodeToString(sum[:]))
}
