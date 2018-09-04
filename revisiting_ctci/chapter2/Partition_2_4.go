/*
take a linked list and partition it around a value

everything smaller than the value is left of all nodes greater than the value
the left and right don't need to be in order

i.e.

3 -> 5 -> 8 -> 5 -> 10 -> 1 -> 2
partition value = 5

1 -> 2 -> 3 -> 5 -> 8 -> 5 -> 10 


unclear:
whether there are instances where things would be mixed up like

 5 -> 3 -> 8 -> 5 

will work on the assumption this is not the case.
check when first answer is completed
*/



package main

import (
	"fmt"
	"./linkedlist"
)

//Remove the node that is currently being viewed from the linked list
func DelNode(n linkedlist.Node) linkedlist.Node {

	next := *n.Next
	prev := *n.Prev

	next.Prev = &prev
	prev.Next = &next

	return next
}


func Partition(ll linkedlist.Node, k int) linkedlist.Node {
	new_ll := linkedlist.Node{Data: k }

	ll = ll.Front()

	if ll.Data < new_ll.Data{
		new_ll.FrontAdd(ll.Data)
	}else{
		new_ll.Add(ll.Data)
	}
	for ll.Next != nil {
		ll = *ll.Next
		if ll.Data < new_ll.Data {
			new_ll.FrontAdd(ll.Data)
		}else{
			new_ll.Add(ll.Data)
		}

	}

	new_ll = DelNode(new_ll)
	return new_ll

}


func main() {
	data := linkedlist.Node{Data: 4}
	data.Add(6)
	data.Add(3)
	data.FrontAdd(12)
	data.Add(2)
	data.Add(2)
	data.FrontAdd(9)
	data.Add(11)
	data.Add(3)
	data.FrontAdd(1)
	data.Add(7)
	data.Add(5)

	fmt.Println(data.Front())

	fmt.Println("Value to partition around: 5")

	data = Partition(data, 5)

	fmt.Println(data.Front())

}

