package main

import(
	"fmt"
)

func fib(length int) []int {
	x2 := 0
	x1 := 1
	var fib_seq = []int{0,1}

	if length == 1 {
		return fib_seq[:1]
	}

	for i := 2; i < length ; i++ {
		new_num := x1 + x2
		fib_seq = append(fib_seq, new_num)
		x2 = x1
		x1 = new_num
	}

	return fib_seq
}

func main(){
	fmt.Println("The first 5 members of the fibonacci sequence are:")
	fmt.Println(fib(5))

	fmt.Println("The first 20 members of the fibonacci sequence are:")
	fmt.Println(fib(20))

	fmt.Println("The first 2 members of the fibonacci sequence are:")
	fmt.Println(fib(2))

	fmt.Println("The first member of the fibonacci sequence is:")
	fmt.Println(fib(1))
}