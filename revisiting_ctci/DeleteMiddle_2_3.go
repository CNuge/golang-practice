/*
given access only to a node in the middle of a linked list,
remove that node from the linked list.

i.e.

1 -> 2 -> 3 -> 4 -> 5
given access to 3rd position
and make the list:
1 -> 2 -> 4 -> 5


*/

package main

import (
	"./linkedlist"
	"fmt"
)

//Remove the node that is currently being viewed from the linked list
func DelNode(n linkedlist.Node) linkedlist.Node {

	next := *n.Next
	prev := *n.Prev

	next.Prev = &prev
	prev.Next = &next

	return next

}

func main() {
	data := linkedlist.Node{Data: 4}

	data.Add(6)
	data.Add(3)
	data.Add(1)
	data.Add(3)
	data.FrontAdd(2)
	data.FrontAdd(2)
	data.FrontAdd(2)

	fmt.Println(data.Front())

	data = DelNode(data)

	fmt.Println(data.Front())

}
