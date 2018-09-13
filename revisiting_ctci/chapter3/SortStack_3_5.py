from stack import Stack



def SortStack(s):
	temp = Stack()

	while s.IsEmpty() == False:
		if temp.IsEmpty():
			temp.Push(s.Pop())

		if s.Peek() > temp.Peek():
			hold = s.Pop()
			s.Push(temp.Pop())
			s.Push(hold)

		elif s.Peek() <= temp.Peek():
			temp.Push(s.Pop())

	return temp


if __name__ == '__main__':

	test = Stack()
	for i in [9,3,1,4,7,5,5,1]:
		test.Push(i)

	print(test)
	print('Sorted low to high:\n')
	print(SortStack(test))		