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


	for len(nodes) > 0 {
		current := *(nodes[0])
		fmt.Println(current)

		nodes = nodes[1:]

		if current.Data == v{
			return true
		}

		nodes = append(nodes, current.Adj...)		

		fmt.Println(nodes)
	}
	return false
}



func main(){

	g1 := graph.Graph{}

	for i:=0 ; i < 5; i++{
		new_node := graph.Node{Data: i*5}
		if len(g1)>0{
			g1[len(g1)-1].Add(new_node)		
		}
		g1.Add(new_node)
	}

	start := g1[0]

	fmt.Println("Graph being used:")
	fmt.Println(g1)

	fmt.Println("Path from 0 -> 19:")
	fmt.Println(IsRoute(start, 19))

	fmt.Println("Path from 0 -> 20:")
	fmt.Println(IsRoute(start, 20))
}