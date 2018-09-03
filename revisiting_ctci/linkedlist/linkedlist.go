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
	next *Node
	prev *Node
}

func (n Node) String() string {
	outstring := fmt.Sprintf("%v", n.data)

	for n.next != nil  {
		outstring = fmt.Sprintf("%v%v", outstring, " -> ")			
		n = *n.next
		outstring = fmt.Sprintf("%v%v", outstring, n.data)			
		}
	return outstring
}


// method to append to a linked list
// take a new number and add it to the linked list
func (n *Node) Add(new_int int) {
	
	new_n := Node{data: new_int}
		
	for n.next != nil {				
		n = n.next
	}
	n.next = &new_n
	new_n.prev = n
}

// method to append to front of a linked list
func (n *Node) FrontAdd(new_int int)  {
	new_n := Node{data: new_int}

	for n.prev != nil {		
		n = n.prev
	}
	n.prev = &new_n
	new_n.next = n

}

