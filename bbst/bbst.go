// 1.) All leaves have count 0
// 2.) Left branch contains <= num, right contains > num

package bst

import (
	"fmt"
	"strings"
)

type Tree struct {
	f float64
	count, layer, depth int
	left, right *Tree
}

func (t *Tree) rotateRight() {
}

func (t *Tree) Add(f float64) (d int) {
	if t.count > 0 { // branch
		if f > t.f {
			d = t.right.Add(f)
		} else {
			d = t.left.Add(f)
		}
		t.depth = d
		return
	}

	// leaf
	if t.f > f {
		t.f, f = f, t.f
	}
	d = t.layer + 1
	t.depth = d
	t.left = &Tree{f:t.f, layer:d}
	t.right = &Tree{f:f, layer:d}
	t.count = 2
	return
}

func (t *Tree) InRange(lo, hi float64) (r []float64) {
	if t.count > 0 { // branch
		if t.f > lo {
			r = t.left.InRange(lo, hi)
		}
		if t.f < hi {
			r = append(r, t.right.InRange(lo, hi)...)
		}
		return
	}

	// leaf
	if t.f < hi {
		r = []float64{t.f}
	}
	return
}


func (t *Tree) Print() {
	space := strings.Repeat(" ", t.layer*2)

	if t.count == 0 {
		fmt.Println(space, int(t.f),",",t.depth)
		return
	}

	t.right.Print()
	fmt.Println(space, int(t.f),",",t.depth)
	t.left.Print()

}
