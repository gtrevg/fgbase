package regexp

import (
	"github.com/vectaport/fgbase"
)

func barFire(n *fgbase.Node) error {
	a := n.Srcs[0]
	b := n.Srcs[1]
	x := n.Dsts[0]
	sink := n.Aux.(bool)

	av := a.SrcGet()
	as := av.(Search).Curr
	ast := av.(Search).State
	bv := b.SrcGet()
	bs := bv.(Search).Curr
	bst := bv.(Search).State

	if ast == Live {
		x.DstPut(av)
		return nil
	}

	if bst == Live {
		x.DstPut(bv)
		return nil
	}

	if ast == Done || bst == Done {
		if sink {
			return nil
		}
		x.DstPut(Search{})
		return nil
	}

	if len(as) > len(bs) {
		x.DstPut(av)
	}
	x.DstPut(bv)
	return nil
}

// FuncBar waits for both inputs and returns the one that matches the shortest string.
// Returns nil if no match at all.
func FuncBar(a, b fgbase.Edge, x fgbase.Edge, sink bool) fgbase.Node {

	node := fgbase.MakeNode("bar", []*fgbase.Edge{&a, &b}, []*fgbase.Edge{&x}, nil, barFire)
	node.Aux = sink
	return node

}
