
import unittest


class Node():
	def __init__(self, Data = None , Next = None, Prev = None):
		self.Data = Data
		self.Next = Next
		self.Prev = Prev

	def __repr__(self):
		return string(self.Data)


class LinkedList():
	def __init__(self):
		self.Node = Node()
	
	def Add(self, num):
		""" add a value to the end of the linked list """
		if self.Node.Data == None:
			self.Node.Data = num
		else:
			current = self.Node
			while current.Next != None:
				current = current.Next
			
			current.Next = Node(num)
			current.Next.Prev = current

	def FrontAdd(self, num):
		""" add a value to the front of the linked list """
		if self.Node.Data == None:
			self.Node.Data = num
		else:
			current = self.Node
			while current.Prev != None:
				current = current.Prev
			
			current.Prev = Node(num)
			current.Prev.Next = current

	def Current(self):
		""" get current location in the lined list """
		return print(self.Node)

	@property
	def Data(self):
		return self.Node.Data

	@property
	def Next(self):
		if self.Node.Next == None:
			return 'At end of the list!'
		return self.Node.Next.Data

	@property
	def Prev(self):
		if self.Node.Prev == None:
			return'At front of the list!'

		return self.Node.Prev.Data

	def MoveForward(self):
		""" advance one position in the linked list """
		if self.Node.Next == None:
			print('At end of the list!')
		else:
			self.Node = self.Node.Next
	
	def MoveBackward(self):
		""" advance one position in the linked list """
		if self.Node.Prev == None:
			print('At front of the list!')
		else:
			self.Node = self.Node.Prev

	def SetFront(self):
		""" set current position to the front of the linked list """
		while self.Node.Prev != None:
			self.Node = self.Node.Prev

	def SetBack(self):
		""" set current position to back of linked list """
		while self.Node.Next != None:
			self.Node = self.Node.Next

	def Front(self):
		""" jump to the front of the linked list """
		front_node = self.Node
		while front_node.Prev != None:
			front_node = front_node.Prev

		return front_node

	def Back(self):
		back_node = self.Node
		while back_node.Next != None:
			back_node = back_node.Next
		
		return back_node

	def __repr__(self):
		self.SetFront()

		outstring = str(self.Node.Data)
		while self.Node.Next != None :

			outstring += " -> "
			self.Node = self.Node.Next
			outstring += str(self.Node.Data)
		
		self.SetFront()
		return outstring

class TestLinkedList(unittest.TestCase):
	@classmethod
	def setUpClass(self):
		"""initiate a linked list and fill with values"""
		self.ll_test = LinkedList()

		self.ll_test.Add(4)
		self.ll_test.Add(5)
		self.ll_test.Add(6)
		self.ll_test.Add(7)
		self.ll_test.FrontAdd(8)
		self.ll_test.FrontAdd(9)

	def test_print(self):
		self.assertEqual(self.ll_test.__repr__(), '9 -> 8 -> 4 -> 5 -> 6 -> 7')
		#make sure its on the right node after printing and not
		#jumping to the front
		self.assertEqual(self.ll_test.Data,  9)


	def test_FrontBack(self):
		#make sure it can call the front and back and 
		#also set the list to the front and the back
		self.ll_test.SetBack()
		self.assertEqual(self.ll_test.Data,  7)

		self.ll_test.SetFront()
		self.assertEqual(self.ll_test.Data,  9)
		
	def test_Move(self):
		self.ll_test.SetFront()
		self.assertEqual(self.ll_test.Next,  8)
		self.assertEqual(self.ll_test.Prev,  'At front of the list!')

		self.ll_test.MoveForward()
		self.assertEqual(self.ll_test.Data,  8)
		self.assertEqual(self.ll_test.Next,  4)
		self.assertEqual(self.ll_test.Prev,  9)

		self.ll_test.SetBack()
		self.ll_test.MoveBackward()

		self.assertEqual(self.ll_test.Data,  6)
		self.assertEqual(self.ll_test.Next,  7)
		self.assertEqual(self.ll_test.Prev,  5)


if __name__ == '__main__':
	unittest.main()
