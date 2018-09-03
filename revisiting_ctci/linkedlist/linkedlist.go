package linkedlist

import (
	"fmt"
)

//node struct
// .data an integer
// .next which references next position (also a node)
// .prev which references previous position (also a node)

type Node struct {
	data int
	next int
	prev int
}

func (n Node) String() string {
	return fmt.Sprintf("%v", n.data)
}


// need a linked list data structure
// composed of nodes which are linked to one another via .next
// the head of the linked list is node 1
type LinkedList []Node

// string method to represent a linked list
// connects the nodes using ->
func (ll LinkedList) String() string {
	outstring := ""
	for e, i := range ll {
		outstring = fmt.Sprintf("%v%v", outstring, i)
		if e != (len(ll) - 1) {
			outstring = fmt.Sprintf("%v%v", outstring, " -> ")
		}
	}
	return fmt.Sprintf("%v", outstring)
}

// method to append to a linked list
// take a new number and add it to the linked list
func (ll *LinkedList) Add(new_int int) {
	
	n := Node{data: new_int}
		
	if len(*ll) > 0 {
		n.prev =  (*ll)[len((*ll))-1].data

		(*ll)[len(*ll)-1].next = n.data
	}

	(*ll) = append((*ll), n)

}

// method to append to front of a linked list
func (ll *LinkedList) FrontAdd(new_int int) {
	n := Node{data: new_int}

	if len(*ll) > 0{
		n.next =  (*ll)[0].data
		(*ll)[0].prev = n.data
	}

	new_ll := LinkedList{n}
	(*ll) = append(new_ll, (*ll)...)

}

