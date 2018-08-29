package flowgraph_test

import (
	"fmt"
	"github.com/vectaport/flowgraph"
	"os"
	"testing"
)

/*=====================================================================*/

func TestMain(m *testing.M) {
	flowgraph.ConfigByFlag(map[string]interface{}{"trace": "Q"})
	os.Exit(m.Run())
}

/*=====================================================================*/

func TestNewEqual(t *testing.T) {

	fmt.Printf("BEGIN:  TestNewEqual\n")

	// Different allocations should not be equal.
	if flowgraph.New("abc") == flowgraph.New("abc") {
		t.Errorf(`New("abc") == New("abc")`)
	}
	if flowgraph.New("abc") == flowgraph.New("xyz") {
		t.Errorf(`New("abc") == New("xyz")`)
	}

	// Same allocation should be equal to itself (not crash).
	g := flowgraph.New("jkl")
	if g != g {
		t.Errorf(`graph != graph`)
	}

	fmt.Printf("END:    TestNewEqual\n")
}

/*=====================================================================*/

type getter struct {
	cnt int
}

func (g *getter) Get(pipe flowgraph.Pipe) (interface{}, error) {
	i := g.cnt
	g.cnt++
	return i, nil
}

func TestInsertIncoming(t *testing.T) {

	fmt.Printf("BEGIN:  TestInsertIncoming")

	fg := flowgraph.New("TestInsertIncoming")
	fg.InsertIncoming("incoming", &getter{})
	fg.InsertSink("sink")

	fg.Run()

	fmt.Printf("END:    TestInsertIncoming\n")
}

/*=====================================================================*/

type putter struct {
	sum int
}

func (p *putter) Put(pipe flowgraph.Pipe, v interface{}) error {
	p.sum += v.(int)
	return nil
}

func TestInsertOutgoing(t *testing.T) {

	fmt.Printf("BEGIN:  TestInsertOutgoing")

	fg := flowgraph.New("TestInsertOutgoing")
	fg.InsertConst("one", 1)
	fg.InsertOutgoing("outgoing", &putter{})

	fg.Run()

	fmt.Printf("END:    TestInsertOutgoing\n")
}

/*=====================================================================*/

type transformer struct {
}

func (t *transformer) Transform(pipe flowgraph.Pipe, v ...interface{}) ([]interface{}, error) {
	xv := v[0].(int) * 2
	return []interface{}{xv}, nil
}

func TestInsertAllOf(t *testing.T) {

	fmt.Printf("BEGIN:  TestInsertTransformer")

	fg := flowgraph.New("TestInsertTransformer")
	fg.InsertConst("one", 1)
	fg.InsertAllOf("double", &transformer{})
	fg.InsertSink("sink")

	fg.Run()

	fmt.Printf("END:    TestInsertTransformer\n")
}

/*=====================================================================*/

func TestInsertArray(t *testing.T) {

	fmt.Printf("BEGIN:  TestInsertArray")

	arr := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fg := flowgraph.New("TestInsertArray")
	fg.InsertArray("array", arr)
	fg.InsertSink("sink")

	fg.Run()

	s := fg.FindPipe("sink").Auxiliary().(flowgraph.SinkStats)
	if s.Cnt != len(arr) {
		t.Fatalf("SinkStats.Cnt != len(arr)\n")
	}
	if s.Sum != 45 {
		t.Fatalf("SinkStats.Sum != sum(arr)\n")
	}

	fmt.Printf("END:    TestInsertArray\n")
}
