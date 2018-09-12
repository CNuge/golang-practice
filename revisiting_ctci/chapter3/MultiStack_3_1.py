"""
use one array to implement three stacks.

i.e.

array := []Stack{3}

have functions to call the array position
and perform the stacks command on the given position
"""


from stack import Stack


class MultiStack():
	def __init__(self):
		self.__stack_list = [Stack(), Stack(), Stack()]

		self.__pos_err = 'MultiStack has three positions avaliable!' + \
								' index: 0 -> 2.'
	def Push(self, dat, pos):
		""" take the given data and add it to the proper stack """
		if pos > 2 :
			raise IndexError(self.__pos_err)

		self.__stack_list[pos].Push(dat)

	def Pop(self, pos):
		if pos > 2 :
			raise IndexError(self.__pos_err)

		return self.__stack_list[pos].Pop()

	def Peek(self, pos):
		if pos > 2 :
			raise IndexError(self.__pos_err)

		return self.__stack_list[pos].Peek()

	def __repr__(self):
		outstring = ''
		for i, x in enumerate(self.__stack_list):
			dat = f'{i}:\t{x}\n'
			outstring += dat
		return outstring

if __name__ == '__main__':

	test = MultiStack()

	test.Push(1,0)	
	test.Push(2,0)	
	test.Push(3,0)	
	test.Push(4,0)	

	test.Push(5,1)	
	test.Push(6,1)	
	test.Push(7,1)	
	test.Push(8,1)	

	test.Push(9,2)	
	test.Push(10,2)	
	test.Push(11,2)	
	test.Push(12,2)

	print(test)
	print('Popped vals:')
	print(test.Pop(0))
	print(test.Pop(1))
	print(test.Pop(2))
	print('\n')

	print(test)