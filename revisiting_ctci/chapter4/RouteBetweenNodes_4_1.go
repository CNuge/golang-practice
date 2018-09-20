/*
Given a directed graph, design an algorithm to find out whther there is a 
route between two nodes
*/

package main

import(
	"./graph"
	"fmt"
)

// given a starting node, see if there is a route to a node
// containing the given value
func IsRoute(n graph.Node, v int) bool{
	nodes := []*graph.Node{&n}

	for len(nodes) > 0{
		current := *(nodes[0])
		nodes = nodes[1:]
		if current.Data == v{
			return true
		}
		nodes = append(nodes, current.Adj...)
	}
	return false
}



func main(){

	g1 := Graph{}

	for i:=0 ;i < 5; i++{
		new:= Node{Data: i*5}
		g1.Add(new)
	}

	start := g1[0]

	fmt.Println("Graph being used:")
	fmt.Println(g1)

	fmt.Println("Path from 0 -> 24:")
	fmt.Println(IsRoute(start, 24))

	fmt.Println("Path from 0 -> 25:")
	fmt.Println(IsRoute(start, 25))
}