/*

implemenet a queue using two stacks.


*/

package main

import(
	"./stack"
	"fmt"
)

type Queue struct{
	s1 stack.Stack
	s2 stack.Stack
}

func (q Queue) String() string{
	return fmt.Sprintf("%v", q.s1)
}

func (q *Queue) Push(i int){
	q.s1.Push(i)
}

func (q *Queue) Pop() int {

	for q.s1.IsEmpty() == false{
		q.s2.Push(q.s1.Pop())
	}
	outval := q.s2.Pop()

	for q.s2.IsEmpty() == false{
		q.s1.Push(q.s2.Pop())
	}

	return outval
}

func (q *Queue) Peek() int{
	for q.s1.IsEmpty() == false{
		q.s2.Push(q.s1.Pop())
	}
	outval := q.s2.Peek()

	for q.s2.IsEmpty() == false{
		q.s1.Push(q.s2.Pop())
	}

	return outval	
}

func (q *Queue) IsEmpty() bool{
	return q.s1.IsEmpty()
}


func main(){


	test_q := Queue{}

	for i := 0; i < 10 ; i++{
		test_q.Push(i)	
	}

	fmt.Println(test_q)
	fmt.Println("Popping from front of queue:")
	fmt.Println(test_q.Pop())
	fmt.Println("Peeking from front of queue:")
	fmt.Println(test_q.Peek())
	fmt.Println("Queue is empty?:")
	fmt.Println(test_q.IsEmpty())
}