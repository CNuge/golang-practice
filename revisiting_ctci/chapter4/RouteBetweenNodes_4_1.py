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