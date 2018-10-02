package graph

// A binary search tree node object
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

	left := BuildBalancedTree(s_list[:middle])
	right := BuildBalancedTree(s_list[middle+1:])

	pos.Left = &left
	pos.Right = &right

	return pos
}
