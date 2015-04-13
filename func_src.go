package flowgraph

import (		
	"bufio"
	"io"
)      			

func srcFire (n *Node) {	 
	x := n.Dsts[0] 		 
	rw := x.Aux.(*bufio.ReadWriter)
	var err error

	// read data string
	x.Val, err = rw.ReadString('\n')
	if err != nil {
		n.Errorf("%v", err)
		close(x.Data)
		x.Data = nil
		return
	}

	// write ack
	_, err = rw.WriteString("\n")
	if err != nil {
		n.Errorf("%v", err)
		close(x.Data)
		x.Data = nil
		return
	}
	rw.Flush()
}

// FuncSrc reads a data value and writes a '\n' acknowledgement.
func FuncSrc(x Edge, rw io.ReadWriter) {

	node := MakeNode("src", nil, []*Edge{&x}, nil, srcFire)
	reader := bufio.NewReader(rw)
	writer := bufio.NewWriter(rw)
	x.Aux = bufio.NewReadWriter(reader, writer)
	node.Run()
	
}
	