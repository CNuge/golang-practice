/*


sort a stack so that the smallest items are on the top.

may use one additional temporary stack but cannot copy elements elsewhere.

using following methods:
-push 
-pop
-peek
-IsEmpty


*/

package main

import(
	"./stack"
	"fmt"
)


func Sort(s stack.Stack) stack.Stack{
	// make a temporary stack
	temp := stack.Stack{}

	// keep looping until everything is out of the original stack
	for s.IsEmpty() == false {
		
		if temp.IsEmpty(){
			// add the top of the input stack to it
			temp.Push(s.Pop())
		}

		// next value is greater than back of temp, need to move
		// the temp top value behind it
		if s.Peek() > temp.Peek(){
			hold := s.Pop()
			s.Push(temp.Pop())
			s.Push(hold)
		// if next up is less than the top of the temp
		// then add it to the top of the stack
		} else if s.Peek() <= temp.Peek(){
			temp.Push(s.Pop())
		}
	}

	return temp
}


func main(){

	test := stack.Stack{}
	test.Push(7)
	test.Push(6)
	test.Push(1)
	test.Push(3)
	test.Push(9)
	test.Push(2)

	fmt.Println(test)

	test = Sort(test)

	fmt.Println(test)

}