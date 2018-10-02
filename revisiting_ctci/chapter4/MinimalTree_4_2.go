/*
Given a sorted array with unique integer elements
write an algorithm to create a binary search tree with
minimal height
*/
package main

import (
	"fmt"
)

type BNode struct{
	Data int
	Left *BNode
	Right *BNode
}



//take a sorted list and construct a binary search tree
func BuildBalancedTree(s_list []int) BNode{
	if len(s_list) == 0{
		return BNode{}
	}
	middle := int(len(s_list)/2)

	pos := BNode{Data : s_list[middle]}


	if len(s_list[:middle]) > 0{
		left := BuildBalancedTree(s_list[:middle])
		pos.Left = &left	
	} else {
		pos.Left = nil
	}
	
	if len(s_list[middle+1:]) > 0 {
		right := BuildBalancedTree(s_list[middle+1:])
		pos.Right = &right
	} else{
		pos.Right = nil
	}

	return pos
}


func main(){
	sorted_array := []int{1,3,5,7,9,11,13}

	ans := BuildBalancedTree(sorted_array)

	fmt.Println(ans.Data == 7)
	left_ans := *ans.Left
	fmt.Println(left_ans.Data == 3)
	right_left_ans := *left_ans.Right
	fmt.Println(right_left_ans.Data == 5)
}
