/*
Given a binary tree design an algorithm which creates a linked list
of all the nodes at each depth
*/
package main

import (
	"../chapter2/linkedlist"

	"fmt"
)

type BNode struct{
	Data int
	Left *BNode
	Right *BNode
}


func BuildBalancedTree(s_list []int) BNode{
	if len(s_list) == 0{
		return BNode{}
	}
	middle := int(len(s_list)/2)

	pos := BNode{Data : s_list[middle]}

	left := BuildBalancedTree(s_list[:middle])
	right := BuildBalancedTree(s_list[middle+1:])

	pos.Left = &left
	pos.Right = &right

	return pos
}

// take the binary tree and turn the layers into a list of linkedlists
// this would let you then traverse the levels horizontally
func DepthLists(in_node BNode) []linkedlist.Node {
	list_of_levels = []linkedlist.Node{}

	level := [in_node.Data]int{}
	
	for len(level) != 0{

	//for each level in the tree, gather all the data

		// for each datum, add it to the linkedlist for the current level

		// advance all nodes to the next level, only add to the next_level
		// if their data is not nil
	}


	return list_of_levels
}


func main(){
	sorted_array := []int{1,3,5,7,9,11,13}

	ans := BuildBalancedTree(sorted_array)
}
