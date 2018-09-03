package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {

	test_ll := Node{data: 1}

	test_ll.Add(1)
	test_ll.Add(3)
	test_ll.Add(8)

	rep_string1 := "1 -> 1 -> 3 -> 8"

	if reflect.DeepEqual(test_ll.String(), rep_string1) != true {
		t.Errorf("String representation not correct! Have:\n%v\nwant:\n%v.", test_ll, rep_string1)

	}

	test_ll.FrontAdd(7)

	rep_string2 := "7 -> 1 -> 1 -> 3 -> 8"

	if reflect.DeepEqual(test_ll.prev.String(), rep_string2) != true {
		t.Errorf("Front add did not perform correctly.\nHave:\n%v\nwant:\n%v.", test_ll, rep_string2)

	}

}
