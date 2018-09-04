/*
find the kth to last element of a singly linked list
*/

package main

import (
	"fmt"
	"./linkedlist"
)



func KthToLast(n linkedlist.Node, k int) linkedlist.Node {
	back := n.Back()

	for i := 1 ; i < k ; i ++ {
		back = *back.Prev
	}
	return back
}


func main(){
	data := linkedlist.Node{Data: 4}
	data.Add(6)
	data.Add(3)
	data.Add(2)
	data.Add(2)
	data.Add(2)
	data.Add(9)
	data.Add(1)
	data.Add(3)

	third_last := KthToLast(data,3)
	fmt.Println(data.Front())
	fmt.Println("The thrid last value is:")
	fmt.Println(third_last.Data)
}