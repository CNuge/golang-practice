/*

check if the data in a linked list is a palindrome

*/

package main

import (
	"./linkedlist"
	"fmt"
)

func LenLL(ll linkedlist.Node) int {
	ll = ll.Front()
	ll_len := 1
	for ll.Next != nil {
		ll = *ll.Next
		ll_len += 1
	}
	return ll_len
}

func IsPalindrome(ll linkedlist.Node) bool {
	ll = ll.Front()
	ll_back := ll.Back()

	ll_len := LenLL(ll)

	for i := 0; i < int(ll_len/2); i++ {
		if ll.Data != ll_back.Data {
			return false
		} else {
			ll = *ll.Next
			ll_back = *ll_back.Prev
		}
	}

	return true
}

func main() {
	data1 := linkedlist.Node{Data: 4}
	data1.Add(6)
	data1.Add(3)
	data1.Add(3)
	data1.Add(6)
	data1.Add(4)
	fmt.Println(data1)
	fmt.Println("Is palindrome:")
	fmt.Println(IsPalindrome(data1))

	data2 := linkedlist.Node{Data: 4}
	data2.Add(6)
	data2.Add(3)
	data2.Add(9)
	data2.Add(3)
	data2.Add(6)
	data2.Add(4)
	fmt.Println(data2)
	fmt.Println("Is palindrome:")
	fmt.Println(IsPalindrome(data2))

	data3 := linkedlist.Node{Data: 4}
	data3.Add(6)
	data3.Add(4)
	data3.Add(9)
	data3.Add(3)
	data3.Add(6)
	data3.Add(4)
	fmt.Println(data3)
	fmt.Println("Is palindrome:")
	fmt.Println(IsPalindrome(data3))

}
