package fast2d

type Point [2]float64

func (p *Point) String() string {
	fmt.Sprintf("{%2.2f, %2.2f}", p[0], p[1])
}

type Node struct {
	point Point
	layer int
	child [2]*Node
}

func MakeTree(points []Point, min, max Point) (arr []Node) {
	arr = make([]Tree, len(points))


	if len(points) > 0 {
		arr[0] = Node{point: points[0]}
		i := 1


		func (n *Node) Add(p Point) {

		}

		for _, v := range points[1:] {

		}

	}
	return
}




