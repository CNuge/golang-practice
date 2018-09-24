package graph

import (
	"fmt"
)

type Node struct {
	Data int
	Adj  []*Node
}

func (n Node) String() string {
	outstring := fmt.Sprintf("%v:", n.Data)
	for _, x := range n.Adj {
		val := *x
		outstring = fmt.Sprintf("%v\t%v", outstring, val.Data)
	}
	return outstring
}

// add to the list of nodes that the current node is connected to
func (n *Node) Add(new_n *Node) {
	n.Adj = append(n.Adj, new_n)
}

type Graph []Node

func (g Graph) String() string {
	outstring := ""
	for _, i := range g {
		outstring = fmt.Sprintf("%v%v\n", outstring, i)
	}
	return outstring
}

func (g *Graph) Add(new_n Node) {
	*g = append(*g, new_n)
}
