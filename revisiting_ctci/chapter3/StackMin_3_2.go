/*

 get the minimum value in a stack

 - efficient is to check on the way in and out
 - need to add a 'times seen' to account for the min getting popped 

*/ 

package main

import (
 	"./stack"
 	"fmt"
 	"errors"
)


func Min(s stack.Stack) (int , error) {
	if s.IsEmpty(){
		return -1, errors.New("The Stack is Empty!") 
	}
	min := s.Ord[0]
	for _, i := range s.Ord{
		if i < min{
			min = i
		}
	}

	return min, nil
}


func main(){

	test := stack.Stack{}
	test.Push(9)
	test.Push(2)
	test.Push(11)
	test.Push(7)
	test.Push(8)

	min, _ := Min(test)
	fmt.Println(min)
}
