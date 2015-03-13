package flowgraph

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

func addfunc(a, b Datum) Datum {
	
	switch a.(type) {
        case int8: { return a.(int8)+b.(int8) }
        case uint8: { return a.(uint8)+b.(uint8) }
        case int16: { return a.(int16)+b.(int16) }
        case uint16: { return a.(uint16)+b.(uint16) }
        case int32: { return a.(int32)+b.(int32) }
        case uint32: { return a.(uint32)+b.(uint32) }
        case int64: { return a.(int64)+b.(int64) }
        case uint64: { return a.(uint64)+b.(uint64) }
	case int: { return a.(int)+b.(int) }
	case uint: { return a.(uint)+b.(uint) }
	case float32: { return a.(float32)+b.(float32) }
	case float64: { return a.(float64)+b.(float64) }
	case complex64: { return a.(complex64)+b.(complex64) }
	case complex128: { return a.(complex128)+b.(complex128) }
	default: { return nil }
	}
}

func AddFunc(a, b, x Conn) {

	pipeid := MakePipe()

	var _a Datum = a.Init_val
	var _b Datum = b.Init_val
	_a_rdy := a.Data_init
	_b_rdy := b.Data_init
	_x_rdy := x.Ack_init

	for {
		fmt.Printf("	add(%d):  _a_rdy,_b_rdy %v,%v  _x_rdy %v\n", pipeid, _a_rdy, _b_rdy, _x_rdy);

		if _a_rdy && _b_rdy && _x_rdy {
			fmt.Println("	add(%d):  writing x and a_req and b_req", pipeid)
			_a_rdy = false
			_b_rdy = false
			_x_rdy = false

			if(reflect.TypeOf(_a)!=reflect.TypeOf(_b)) {
				_,nm,ln,_ := runtime.Caller(0)
				x.Data <-  errors.New(fmt.Sprintf("%s:%d (pipeid %d)  type mismatch (%v,%v)", nm, ln, pipeid, reflect.TypeOf(_a), reflect.TypeOf(_b)))
			} else {
				x.Data <- addfunc(_a, _b)
			}

			a.Ack <- true
			b.Ack <- true
			fmt.Println("	add(%d):  done writing x and a_req and b_req", pipeid)
		}

		fmt.Println("	add(%d):  select")
		select {
		case _a = <-a.Data:
			{
				fmt.Printf("	add(%d):  a read %v --  %v\n", pipeid, reflect.TypeOf(_a), _a)
				_a_rdy = true
			}
		case _b = <-b.Data:
			{
				fmt.Printf("	add(%d):  b read %v --  %v\n", pipeid, reflect.TypeOf(_b), _b)
				_b_rdy = true
			}
		case _x_rdy = <-x.Ack:
			fmt.Println("	add(%d):  x_req read", pipeid)
		}

	}

}
