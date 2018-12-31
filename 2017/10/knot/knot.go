package knot

type knot struct {
	String [256]int
	Pos    int
	Skip   int
}

func Sum(v []byte) [16]byte {
	k := newKnot()
	v = append(v, []byte{17, 31, 73, 47, 23}...)
	for i := 0; i < 64; i++ {
		for _, l := range v {
			k.apply(l)
		}
	}
	return k.dense()
}

func newKnot() *knot {
	var k knot
	for i := 0; i < len(k.String); i++ {
		k.String[i] = i
	}
	return &k
}

func (k *knot) apply(length byte) {
	if length > 1 {
		i, j := k.Pos, (k.Pos+int(length)-1)%len(k.String)
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

	k.Pos += int(length) + k.Skip
	k.Pos %= len(k.String)
	k.Skip++
}

func (k *knot) dense() (o [16]byte) {
	for b := 0; b < 16; b++ {
		o[b] = byte(k.String[b*16])
		for i := 1; i < 16; i++ {
			o[b] ^= byte(k.String[b*16+i])
		}
	}
	return o
}
