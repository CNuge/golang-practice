/*
remove duplicate values from an unsorted linked list
*/

package main

import (
	"./linkedlist"
	"fmt"
)

func DeDup(n linkedlist.Node) linkedlist.Node {
	new_ll := linkedlist.Node{Data: n.Data}

	seen_vals := []int{}

	for n.Next != nil {
		n = *n.Next
		seen := 0
		for _, val := range seen_vals {
			if n.Data == val {
				seen = 1
			}
		}
		if seen != 1 {
			seen_vals = append(seen_vals, n.Data)
			new_ll.Add(n.Data)
		}
	}

	return new_ll
}

func main() {

	data := linkedlist.Node{Data: 4}
	data.Add(6)
	data.Add(3)
	data.Add(2)
	data.Add(2)
	data.Add(2)
	data.Add(1)
	data.Add(3)

	data = DeDup(data)

	fmt.Println(data)
	fmt.Println("4 -> 6 -> 3 -> 2 -> 1")

}
