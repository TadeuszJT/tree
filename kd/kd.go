package kd

import (
	"fmt"
	"strings"
)

const k = 2

type Point [k]float64

func (p Point) String() string {
	return fmt.Sprintf("{%2.2f, %2.2f}", p[0], p[1])
}

type KDTree struct {
	P           Point
	layer       int
	left, right *KDTree
}

func (t *KDTree) Add(p Point) {
	cd := t.layer % k

	if p[cd] < t.P[cd] {
		if t.left != nil {
			t.left.Add(p)
		} else {
			t.left = &KDTree{P: p, layer: t.layer + 1}
		}
	} else {
		if t.right != nil {
			t.right.Add(p)
		} else {
			t.right = &KDTree{P: p, layer: t.layer + 1}
		}
	}
}

func Naive (p []Point, min, max Point) (r []Point) {
	for _, v := range p {
		in := true
		for i := 0; i < k; i++ {
			if v[i] < min[i] || v[i] > max[i] {
				in = false
				break
			}
		}
		if in {
			r = append(r, v)
		}
	}
	return
}

func (t *KDTree) InRange(min, max Point) (r []Point) {
	// test point

	if t.P[0] > min[0] && t.P[1] > min[1] && t.P[0] < max[0] && t.P[1] < max[1] {
		r = []Point{t.P}
	}

	//inRange := true
	//for i := 0; i < k; i++ {
		//if t.P[i] < min[i] || t.P[i] > max[i] {
			//inRange = false
			//break
		//}
	//}
	//if inRange {
		//r = []Point{t.P}
	//}

	// return branches
	cd := t.layer % k
	if t.left != nil && t.P[cd] > min[cd] {
		r = append(r, t.left.InRange(min, max)...)
	}
	if t.right != nil && t.P[cd] < max[cd] {
		r = append(r, t.right.InRange(min, max)...)
	}
	return
}

func (t *KDTree) Print() {
	space := strings.Repeat(" ", t.layer*2)
	if t.right != nil {
		t.right.Print()
	}
	fmt.Println(space, t.P)
	if t.left != nil {
		t.left.Print()
	}
}
