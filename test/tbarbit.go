package main

import (
	"flowgraph"
	"fmt"
	"reflect"
	"time"
)

func tbi(a flowgraph.Conn) {

	pipeid:=flowgraph.MakePipe()

	var _a flowgraph.Datum = 0
	_a_rdy := a.Ack_init

	for {


		if _a_rdy {
			fmt.Printf("tbi(%d):  writing a.Data: %d\n", pipeid, _a.(int))
			_a_rdy = false
			a.Data <- _a
			_a = (_a.(int) + 1)
		}

		fmt.Printf("tbi(%d):  select", pipeid)
		select {
		case _a_rdy = <-a.Ack:
			fmt.Printf("tbi(%d):  a.Ack read\n", pipeid)
			
			
		}
	}
	
}

func tbo(x flowgraph.Conn) {
	
	pipeid:=flowgraph.MakePipe()
	
	var _x flowgraph.Datum
	_x_rdy := x.Data_init

	for {
		// fmt.Println("		tbo:  _x_rdy", _x_rdy)
		if _x_rdy {
			fmt.Printf("		tbo(%d):  writing x.Ack\n", pipeid)
			x.Ack <- true
			_x_rdy = false
		}

		fmt.Println("		tbo:  select")
		select {
		case _x = <-x.Data:
			{
				fmt.Printf("		tbo(%d):  x read %v --  %v\n", pipeid, reflect.TypeOf(_x), _x)
				_x_rdy = true
			}
		}

	}

}

func main() {

	a := flowgraph.MakeConn(false,true,nil)
	b := flowgraph.MakeConn(false,true,nil)
	x := flowgraph.MakeConn(false,true,nil)

	go tbi(a)
	go tbi(b)
	go flowgraph.ArbitFunc(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}

