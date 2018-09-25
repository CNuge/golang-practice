"""
Given a sorted array with unique integer elements
write an algorithm to create a binary dearch tree with
minimal height
"""
from graph import Node, Graph

class BalancedNode():
	def __init__(self, data=None):
		#this stores the data found at the given node
		self.__data = data
		#this stores the adjacent edges in the graph, that can
		#be reached from the given position
		self.__left = None
		self.__right = None

	def LeftAdd(self, data):
		"""take a new node and add it to the list of adjacent positions """
		self.__left = data

	def RightAdd(self, data):
		"""take a new node and add it to the list of adjacent positions """
		self.__right = data

	def ChangeVal(self, value):
		""" change the value stored for the given node """
		self.__data = value

	@property
	def Data(self):
		return self.__data

	@property
	def Left(self):
		return self.__left

	@property
	def Right(self):
		return self.__right

	def __repr__(self):
		return f'{self.Data}'


def BuildBalancedGraph(sorted_array):
	if len(sorted_array) == 0:
		return []

	middle = int(len(sorted_array)/2)
	pos = BalancedNode(sorted_array[middle])

	left = BuildBalancedGraph(sorted_array[:middle])
	right = BuildBalancedGraph(sorted_array[middle+1:])

	pos.LeftAdd(left)
	pos.RightAdd(right)

	return pos

if __name__ == '__main__':


	sorted_array = [1,3,5,7,9,11,13]

	ans = BuildBalancedGraph(sorted_array)

	ans.Data
	x = ans.Left
	y = x.Left
	y.Left