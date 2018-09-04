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

import(
	"fmt"
	"./linkedlist"
)


// take a Node and scroll through the linked list
// to return the front position


func DelNode(n Node) Node {

}



func main(){
	data := linkedlist.Node{Data : 4}

	data.Add(6)
	data.Add(3)
	data.Add(1)
	data.Add(3)
	data.FrontAdd(2)
	data.FrontAdd(2)
	data.FrontAdd(2)

	fmt.Println(data.Front())

}