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


	if s.Left != nil {

		left := *s.Left
	
		fmt.Printf("data: %v , left: %v\n",s.Data, left.Data)

		if left.Data > s.Data {
			return false
		}

		left_ans = ValidateBST(left)
	} else {
		left_ans = true
	}

	if  s.Right != nil {
		right := *s.Right	
		
		fmt.Printf("data: %v , right: %v\n",s.Data, right.Data)


		if right.Data < s.Data{
			return false
		}

		right_ans = ValidateBST(right)		
	} else {
		right_ans = true
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