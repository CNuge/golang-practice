"""
Given a directed graph, design an algorithm to find out whther there is a 
route between two nodes
"""

from graph import Node, Graph


def IsRoute(TopNode, dest):
	""" conduct a breadth first search to see if there is a path between
		the start node and the destination"""
	nodes = [TopNode]
	while len(nodes) > 0 :
		current = nodes.pop(0)
		if current.Data == dest:
			return True
		nodes.extend(current.adj)

	return False


#conduct a breadth first search, beginning with a seed node and visiting
#all of its progeny, and their progeny in turn.




if __name__ == '__main__':

	#initiate a graph
	test_nodes = Graph()

	for i in range(6):
		x = i*3
		new = Node(x)

		if len(test_nodes.Nodes) > 0:
			test_nodes.Nodes[-1].Add(new)
		test_nodes.Add(new)


	IsRoute(test_nodes.Nodes[0], 14) #False

	IsRoute(test_nodes.Nodes[0], 12) #True
