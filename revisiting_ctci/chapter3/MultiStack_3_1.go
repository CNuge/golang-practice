
/*
use one array to implement three stacks.

i.e.

array := []Stack{3}

have functions to call the array position 
and perform the stacks command on the given position

*/


package main


import(
	"fmt"
	"./stack"
)


type MultiStack [3]stack.Stack

func (m MultiStack) String() string{
	outstring := ""
	for i, dat := range m {
		outstring = fmt.Sprintf("%v%v\t%v\n",outstring, i, dat )
	}
	return outstring
}

func (m MultiStack) Push(pos, dat int){
	m[pos].Push(dat)
}

func (m MultiStack) Pop(pos int) int {
	output :=  m[pos].Pop()
	return output
}

func (m MultiStack) IsEmpty(pos int) bool{
	if len(m[pos].Ord) == 0 {
		return true
	}
	return false
}


func main(){

	test := MultiStack{}
	test.Push(1,0)
	test.Push(2,0)
	test.Push(3,0)
	test.Push(4,0)
	fmt.Println(test)

	test.Pop(0)

	test.Push(4,1)
	test.Push(5,1)
	test.Push(6,1)

	test.Push(7,2)
	test.Push(8,2)
	test.Push(9,2)

	fmt.Println(test)

}