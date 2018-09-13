from stack import Stack

class SetOfStacks():
	def __init__(self, max_len):
		self.__all = [Stack()]
		self.__max_len = max_len

	def Push(self, data):
		if len(self.__all[-1].Data) >= self.__max_len:
			new = Stack()
			new.Push(data)
			self.__all.append(new)
		else:
			self.__all[-1].Push(data)

	def Pop(self):
		out = self.__all[-1].Pop()
		if len(self.__all[-1].Data) == 0:
			self.__all = self.__all[:-1]
		return out

	def Peek(self):
		return self.__all[-1].Peek()

	def __repr__(self):
		outstring = ''
		for i, data in enumerate(self.__all):
			line = f'{i}:\t{data}\n'
			outstring += line
		return outstring


if __name__ == '__main__':

	test = SetOfStacks(5)

	for i in range(17):
		test.Push(i)

	print(test)
	
	print('Popping:\n')
	for i in range(3):
		print(test.Pop())

	print('\nFinal:\n')
	print(test)