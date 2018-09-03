package linkedlist

import (
	"fmt"
)

//node struct
// .Data an integer
// .Next which references.Next position (also a node)
// .Prev which references.Previous position (also a node)

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func (n Node) String() string {
	outstring := fmt.Sprintf("%v", n.Data)

	for n.Next != nil  {
		outstring = fmt.Sprintf("%v%v", outstring, " -> ")			
		n = *n.Next
		outstring = fmt.Sprintf("%v%v", outstring, n.Data)			
		}
	return outstring
}


// method to append to a linked list
// take a new number and add it to the linked list
func (n *Node) Add(new_int int) {
	
	new_n := Node{Data: new_int}
		
	for n.Next != nil {				
		n = n.Next
	}
	n.Next = &new_n
	new_n.Prev = n
}

// method to append to front of a linked list
func (n *Node) FrontAdd(new_int int)  {
	new_n := Node{Data: new_int}

	for n.Prev != nil {		
		n = n.Prev
	}
	n.Prev = &new_n
	new_n.Next = n

}

