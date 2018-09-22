package range2d

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("{%1.2f, %1.2f}", p.X, p.Y)
}

type YTree struct {
	p Point
	count, layer int
	left, right *YTree
}

type Tree struct {
	p Point
	count, layer int
	left, right *Tree
	ytree *YTree
}

func (t *YTree) Add(p Point) {
	if t.count > 0 { // branch
		if p.Y > t.p.Y {
			t.right.Add(p)
		} else {
			t.left.Add(p)
		}
		t.count++
		return
	}

	// leaf
	if t.p.Y > p.Y {
		t.p, p = p, t.p
	}
	t.left = &YTree{p:t.p, layer:t.layer + 1}
	t.right = &YTree{p:p, layer: t.layer + 1}
	t.count = 2
}

func (t *YTree) Print() {
	space := strings.Repeat(" ", t.layer*4)

	if t.count > 0 {
		t.right.Print()
		fmt.Println(space, t.p)
		t.left.Print()
		return
	}

	fmt.Println(space, t.p)
}

func (t *Tree) Add(p Point) {
	if t.count > 0 { // branch
		t.ytree.Add(p)
		if p.X > t.p.X {
			t.right.Add(p)
		} else {
			t.left.Add(p)
		}
		t.count++
		return
	}

	// leaf
	if t.p.X > p.X {
		t.p, p = p, t.p
	}
	t.left = &Tree{p:t.p, layer:t.layer + 1}
	t.right = &Tree{p:p, layer:t.layer + 1}
	t.ytree = &YTree{p:t.p}
	t.ytree.Add(p)
	t.count = 2
}

func (t *YTree) InRange(xmin, xmax, ymin, ymax float64) (r []Point) {
	if t.count > 0 { // branch
		if t.p.Y > ymin {
			r = t.left.InRange(xmin, xmax, ymin, ymax)
		}
		if t.p.Y < ymax {
			r = append(r, t.right.InRange(xmin, xmax, ymin, ymax)...)
		}
		return
	}

	// leaf
	if t.p.Y < ymax && t.p.Y > ymin && t.p.X < xmax && t.p.X > xmin {
		r = []Point{t.p}
	}
	return
}

func (t *Tree) InRange(xmin, xmax, ymin, ymax float64) []Point {
	if t.count > 0 { // branch
		if xmin > t.p.X { // go right
			return t.right.InRange(xmin, xmax, ymin, ymax)
		} else if xmax < t.p.X { // go left
			return t.left.InRange(xmin, xmax, ymin, ymax)
		} else {
			// search y
			t.ytree.Print()
			return t.ytree.InRange(xmin, xmax, ymin, ymax)
		}
	}

	// leaf
	if t.p.X >= xmin && t.p.X < xmax && t.p.Y >= ymin && t.p.Y < ymax {
		return []Point{t.p}
	}
	return []Point{}
}

func (t *Tree) Print() {
	space := strings.Repeat(" ", t.layer*4)

	if t.count > 0 {
		t.right.Print()
		fmt.Println(space, t.p)
		t.left.Print()
		return
	}

	fmt.Println(space, t.p)
}
