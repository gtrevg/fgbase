package flowgraph

import (
)

func strcnd_func (n *Node) {
	a := n.Srcs[0]
	x := n.Dsts[0]
	y := n.Dsts[1]
	x.Val = nil
	y.Val = nil
	if (ZeroTest(a.Val)) {
		x.Val = a.Val
	} else {
		y.Val = a.Val
	}
}

func strcnd_rdy (n *Node) bool {
	a := n.Srcs[0]
	x := n.Dsts[0]
	y := n.Dsts[1]
	if a.Rdy {
		if ZeroTest(a.Val) {
			return x.Rdy
		} else {
			return y.Rdy
		}
	} else {
		return false
	}
}

// Steer condition goroutine
func FuncStrCnd(a, x, y Edge) {

	node := MakeNode2("strcnd", []*Edge{&a}, []*Edge{&x, &y}, strcnd_rdy, strcnd_func)
	node.Run()

}
