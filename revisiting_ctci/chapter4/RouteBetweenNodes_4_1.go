/*
Given a directed graph, design an algorithm to find out whther there is a
route between two nodes
*/

package main

import (
	"./graph"
	"fmt"
)

// given a starting node, see if there is a route to a node
// containing the given value
func IsRoute(n graph.Node, v int) bool {
	nodes := []*graph.Node{&n}

	for len(nodes) > 0 {
		current_p := nodes[0]
		current := *current_p

		nodes = nodes[1:]

		if current.Data == v {
			return true
		}

		nodes = append(nodes, current.Adj...)
	}
	return false
}

func main() {

	n1 := graph.Node{Data: 0}
	n2 := graph.Node{Data: 5}
	n3 := graph.Node{Data: 10}
	n4 := graph.Node{Data: 15}
	n5 := graph.Node{Data: 20}

	n1.Add(&n2)
	n2.Add(&n3)
	n3.Add(&n4)
	n4.Add(&n5)

	g1 := graph.Graph{n1, n2, n3, n4, n5}

	start := g1[0]

	fmt.Println(start)
	fmt.Println(*(start.Adj[0]))

	fmt.Println("Graph being used:")
	fmt.Println(g1)

	fmt.Println("Path from 0 -> 19:")
	fmt.Println(IsRoute(start, 19))

	fmt.Println("Path from 0 -> 20:")
	fmt.Println(IsRoute(start, 20))
}
