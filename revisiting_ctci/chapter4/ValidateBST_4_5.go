/*
implement a function to check if a binary tree is a binary search tree
*/

package main

import(
	"./graph"
	"fmt"
)



func ValidateBST(s graph.BNode) bool {

	var left_ans bool
	var right_ans bool

	if s.Left == nil{
		left_ans = true
	} else {

		left := *s.Left
		if left.Data > s.Data{
			return false
		}

		left_ans = ValidateBST(left)
	}

	if s.Right == nil {
		right_ans = true	
	} else{
		right := *s.Right	
		if right.Data < s.Data{
			return false
		}

		right_ans = ValidateBST(right)		
	}

	if right_ans == true && left_ans == true {
		return true
	}

	return false
}


func main(){

	sorted_array := []int{1,3,5,7,9,11,13}

	test_tree_1 := graph.BuildBalancedTree(sorted_array)

	fmt.Println(sorted_array)
	fmt.Println(ValidateBST(test_tree_1))

	non_sorted_array := []int{1,7,9,11,3,5,13}

	test_tree_2 := graph.BuildBalancedTree(non_sorted_array)

	fmt.Println(non_sorted_array)
	fmt.Println(ValidateBST(test_tree_2))

}