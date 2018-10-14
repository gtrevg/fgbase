package fgbase

import ()

func SteervRdy(n *Node) bool {
	if len(n.Srcs) == 1 {
		return SteercRdy(n)
	}

	a := n.Srcs[0]
	b := n.Srcs[1]
	if a.SrcRdy(n) && b.SrcRdy(n) {
		if ZeroTest(a.Val) {
			return n.Dsts[0].DstRdy(n)
		}
		i := min(Int(a.Val), len(n.Dsts)-1)
		return n.Dsts[i].DstRdy(n)
	}
	return false
}

func SteervFire(n *Node) error {
	if len(n.Srcs) == 1 {
		return SteercFire(n)
	}

	a := n.Srcs[0]
	b := n.Srcs[1]
	av := a.SrcGet()
	if ZeroTest(av) {
		n.Dsts[0].DstPut(b.SrcGet())
	} else {
		i := min(Int(av), len(n.Dsts)-1)
		n.Dsts[i].DstPut(b.SrcGet())
	}
	return nil
}

// FuncSteerv steers the second value by the first (if a==0 { x = b } else { y = b }).
func FuncSteerv(a, b, x, y Edge) Node {

	node := MakeNode("steerv", []*Edge{&a, &b}, []*Edge{&x, &y}, SteervRdy, SteervFire)
	return node

}
