
import unittest


class Node():
	def __init__(self, Data = None , Next = None, Prev = None):
		self.Data = data
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
		current = self.Node
		while current.Next != None:
			current = current.Next
		else:
			current.Next = Node(num)

	def FrontAdd(self, num):
		""" add a value to the front of the linked list """
		if self.Node.Data == None:
			self.Node.Data = num
		current = self.Node
		while current.Prev != None:
			current = current.Prev
		else:
			current.Prev = Node(num)

	def Current(self):
		""" get current location in the lined list """
		return print(self.Node)

	@property
	def Next(self):
		if self.Node.Next == None:
			print('At end of the list!')
		return string(self.Node.Next.Data)

	@property
	def Prev(self):
		if self.Node.Prev == None:
			print('At front of the list!')

		return string(self.Node.Prev.Data)

	def MoveForward(self):
		""" advance one position in the linked list """
		if self.Node.Next == None:
			print('At end of the list!')
		else:
			self.Node = self.Node.Next
	
	def MoveBackwards(self):
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
		to_show = self.Front()

		outstring = string(to_show.Data)
		while to_show.Next != None :
			outstring += " -> "
			to_show = to_show.Next
			outstring += string(to_show.Data)

		return outstring

class TestLinkedList(unittest.TestCase):
	@classmethod
	def setUpClass(self):
		"""initiate a linked list and fill with values"""

	def test_print(self):

		#make sure its on the right node after printing and not
		#jumping to the front

	def test_FrontBack(self):
		#make sure it can call the front and back and 
		#also set the list to the front and the back
		
	def test_print(self):

if __name__ == '__main__':
	unittest.main()
