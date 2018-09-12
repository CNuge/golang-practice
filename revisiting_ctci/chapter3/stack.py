import unittest


class Stack():
	def __init__(self):
		self.__dat = []

	def Push(self, n):
		""" add a number to the top of the stack """
		self.__dat.append(n)

	def Peek(self):
		""" look at the number at the top of the stack """
		return self.__dat[-1]

	def Pop(self):
		return self.__dat.pop(-1)

	def __repr__(self):
		outstring = ''
		""" show the data in the stack """
		for i in reversed(self.__dat):
			outstring += str(i)
			outstring += '\t'
		return outstring


class TestStack(unittest.TestCase):
	@classmethod
	def setUpClass(self):
		self.stackCase = Stack()

		self.stackCase.Push(3)
		self.stackCase.Push(6)
		self.stackCase.Push(7)
		self.stackCase.Push(2)
		self.stackCase.Push(4)

	def test_peek(self):
		self.assertEqual(self.stackCase.Peek(), 4)

	def test_pop(self):
		self.assertEqual(self.stackCase.Pop(), 4)
		self.assertEqual(self.stackCase.Peek(), 2)

	def test_print(self):
		self.assertEqual(self.stackCase.__repr__(), '2\t7\t6\t3\t')

if __name__ == '__main__':
	unittest.main()
