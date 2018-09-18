import unittest

class Node():
	def __init__(self, data=None):
		#this stores the data found at the given node
		self.__data = data
		#this stores the adjacent edges in the graph, that can
		#be reached from the given position
		self.adj = []

	def Add(self, *args):
		"""take a new node and add it to the list of adjacent positions """
		for node in args:
			self.adj.append(node)

	def ChangeVal(self, value):
		""" change the value stored for the given node """
		self.__data = value

	@property
	def Data(self):
		return self.__data

	@property
	def Children(self):
		if len(self.adj) == 0:
			return ''
		out = [str(x.Data) for x in self.adj]
		return '\t'.join(out)

	def __repr__(self):
		return f'{self.Data}:\t{self.Children}'


class Graph():
	def __init__(self):
		self.Nodes = []

	def Add(self, *args):
		for node in args:
			self.Nodes.append(node)

	def __repr__(self):
		outstring = ''
		for n in self.Nodes:
			outstring += str(n)
			outstring += '\n'
		return outstring


class TestNodes(unittest.TestCase):
	@classmethod
	def setUpClass(self):
		self.n1 = Node(1)
		self.n2 = Node(2)
		self.n3 = Node(3)
		self.n4 = Node(4)

		self.n1.Add(self.n2)
		self.n2.Add(self.n3,self.n4)

		self.g_test = Graph()
		self.g_test.Add(self.n1,self.n2,self.n3,self.n4)

	def test_NodeStrings(self):
		self.assertEqual(self.n1.Children , '2')
		self.assertEqual(self.n1.Data , 1 )
		self.assertEqual(str(self.n1) , '1:\t2' )
		self.assertEqual(self.n2.Children ,'3\t4' )
		self.assertEqual(str(self.n2) , '2:\t3\t4' )
		self.assertEqual(self.n2.adj, [self.n3, self.n4])
		
	def test_GraphStr(self):
		self.assertEqual(str(self.g_test) , '1:\t2\n2:\t3\t4\n3:\t\n4:\t\n')

if __name__ == '__main__':
	unittest.main()