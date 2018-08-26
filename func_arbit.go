package flowgraph

import ()

func arbitFire(n *Node) {
	a := n.Srcs[0]
	b := n.Srcs[1]
	x := n.Dsts[0]
	if a.SrcRdy(n) && (!b.SrcRdy(n) || !n.Aux.(bool)) {
		n.Aux = true // aLast = true
		x.DstPut(a.SrcGet())
		if b.SrcRdy(n) {
		}
	} else if b.SrcRdy(n) {
		n.Aux = false // aLast = false
		x.DstPut(b.SrcGet())
		if a.SrcRdy(n) {
		}
	}
}

func arbitRdy(n *Node) bool {
	return (n.Srcs[0].SrcRdy(n) || n.Srcs[1].SrcRdy(n)) && n.Dsts[0].DstRdy(n)
}

// FuncArbit arbitrates between two values (select { case a: x = a case b: x = b }).
func FuncArbit(a, b, x Edge) Node {

	node := MakeNode("arbit", []*Edge{&a, &b}, []*Edge{&x}, arbitRdy, arbitFire)
	node.Aux = false // aux value that means "a" won the arbitration last
	return node

}
