package linkedlist


import(
	"fmt"
	
)


//node struct
// .data an integer
// .next which references next position (also a node)
// .prev which references previous position (also a node)

type Node struct{
	data int
	next int
	prev int
}

// need a linked list data structure
// composed of nodes which are linked to one another via .next
// the head of the linked list is node 1
type LinkedList []Node


// string method to represent a linked list 
// connects the nodes using ->
func (ll LinkedList) String() string{
	outstring := ""  
	for e , i := range ll {
		outstring += string(i.data)
		if e != (len(ll) - 1) {
			outstring += " -> "
		}
	}
}

// method to append to a linked list
// take a new number and add it to the linked list
func (ll *LinkedList) Add(new_int int) {
	n := Node{data : new_int
				prev : ll[len(ll)-1]}

	ll[len(ll)-1].next = n

	ll = append(ll, n)

}


// method to append to front of a linked list
func (ll *LinkedList) FrontAdd(new_int int){
	n := Node{data : new_int
			next : ll[0]}

	ll[0].prev = n

	new_ll := LinkedList{n}
	ll = append(new_ll, ll...)

}


