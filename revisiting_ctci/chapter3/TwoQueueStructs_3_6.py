"""

animal shelter holds dogs and cats.

first in shelter must be first adopted (queue)

people can select whether they want a dog or a cat,
otherwise if no choice they get the oldest.


implement this setup


"""



class Queue():
	def __init__(self):
		self.__dat = []

	def Push(self, n):
		""" add a number to the top of the stack """
		self.__dat.append(n)

	def Peek(self):
		""" look at the number at the top of the stack """
		return self.__dat[0]

	def Pop(self):
		return self.__dat.pop(0)

	def IsEmpty(self):
		if len(self.__dat) == 0:
			return True
		return False

	@property
	def Data(self):
		return [x for x in self.__dat]

	def __repr__(self):
		outstring = ''
		""" show the data in the stack """
		for i in reversed(self.__dat):
			outstring += str(i)
			outstring += '\t'
		return outstring
