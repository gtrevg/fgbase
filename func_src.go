package fgbase

import (
	"bufio"
	"io"
)

func srcFire(n *Node) {
	x := n.Dsts[0]
	rw := n.Aux.(*bufio.ReadWriter)

	// read data string
	xv, err := rw.ReadString('\n')
	if err != nil {
		n.LogError("%v", err)
		x.CloseData()
		return
	}
	x.DstPut(xv)

	// write ack
	_, err = rw.WriteString("\n")
	if err != nil {
		n.LogError("%v", err)
		x.CloseData()
		return
	}
	rw.Flush()
}

// FuncSrc reads a data value and writes a '\n' acknowledgement.
func FuncSrc(x Edge, rw io.ReadWriter) Node {

	node := MakeNode("src", nil, []*Edge{&x}, nil, srcFire)
	reader := bufio.NewReader(rw)
	writer := bufio.NewWriter(rw)
	node.Aux = bufio.NewReadWriter(reader, writer)
	return node

}
