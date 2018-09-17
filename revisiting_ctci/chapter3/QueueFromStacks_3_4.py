import unittest
from stack import Stack

class Queue():
	def __init__(self):
		self.s1 = Stack()
		self.s2 = Stack()

	def Push(self, i):
		self.s1.Push(i)

	def Pop(self):
		while self.s1.IsEmpty() == False:
			self.s2.Push(self.s1.Pop())
		outdat = self.s2.Pop()
		while self.s2.IsEmpty() == False:
			self.s1.Push(self.s2.Pop())
		return outdat

	def Peek(self):
		while self.s1.IsEmpty() == False:
			self.s2.Push(self.s1.Pop())
		outdat = self.s2.Peek()
		while self.s2.IsEmpty() == False:
			self.s1.Push(self.s2.Pop())
		return outdat


	def IsEmpty(self):
		return self.s1.IsEmpty()

	def __repr__(self):
		return str(self.s1)[::-1] 



class TestStack(unittest.TestCase):
	@classmethod
	def setUpClass(self):
		self.stackCase = Queue()

		self.stackCase.Push(3)
		self.stackCase.Push(6)
		self.stackCase.Push(7)
		self.stackCase.Push(2)
		self.stackCase.Push(4)

	def test_peek(self):
		self.assertEqual(self.stackCase.Peek(), 3)

	def test_pop(self):
		self.assertEqual(self.stackCase.Pop(), 3)
		self.assertEqual(self.stackCase.Peek(), 6)

	def test_print(self):
		self.assertEqual(self.stackCase.__repr__(), '6\t7\t2\t4')


if __name__ == '__main__':
	unittest.main()
