package graph

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	n1 := Node{Data: 7}

	n2 := Node{Data: 1}
	n3 := Node{Data: 2}

	n1.Add(n2)
	n1.Add(n3)

	n1_ex_string := "7:\t1\t2"

	if reflect.DeepEqual(n1.String(), n1_ex_string) != true {
		t.Errorf("String representation of Node not correct! Have:\n%v\nwant:\n%v",
			n1, n1_ex_string)
	}
}

func TestGraph(t *testing.T) {
	n1 := Node{Data: 7}
	n2 := Node{Data: 1}
	n3 := Node{Data: 2}
	n4 := Node{Data: 8}
	n5 := Node{Data: 9}

	n1.Add(n2)
	n2.Add(n3)
	n2.Add(n4)
	n4.Add(n5)
	n5.Add(n1)

	g1 := Graph{n1, n2, n3, n4, n5}

	g_string := "7:\t1\n" +
		"1:\t2\t8\n" +
		"2:\n" +
		"8:\t9\n" +
		"9:\t7\n"

	if reflect.DeepEqual(g1.String(), g_string) != true {
		t.Errorf("String representation of Graph not correct! Have:\n%v\nwant:\n%v",
			g1, g_string)
	}
}
