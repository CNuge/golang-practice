

package linkedlist


import(
	"fmt"
	"reflect"
	"testing"
)



func TestLinkedList(t *testing.T){

	var test_ll LinkedList

	test_ll.add(1)
	test_ll.add(1)
	test_ll.add(3)
	test_ll.add(8)

	rep_string := "1 -> 1 -> 3 -> 8"

	if reflect.DeepEqual(fmt.Printf(test_ll) , ) != true {
		t.Errorf("String representation not correct! Have:\n%v\nwant:\n%v.", test_ll, rep_string)

	}
 

}