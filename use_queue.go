package main

import(
	"fmt"
	"./queue"
)

/*
this shows how to implement a package, and also acts as a test of the code I wrote

So I have built the package queue, which can be accessed via an import from ./queue
note the package queue in the containing files matches the directory name.

All of the functions I deined in that package are then accessiable from this external
location, but I must call them as queue.Queue... as is convention.

These lines have been turned into a unit test witin the package

This was helpful:
https://blog.golang.org/package-names
*/


func main(){
	q := queue.Queue{}
	q.Add(7)
	q.Add(8)
	q.Add(9)
	fmt.Println(q)
	fmt.Println(q.Peek())
	fmt.Println(q)
	fmt.Println(q.Remove())
	fmt.Println(q.Peek())
	fmt.Println(q)
}