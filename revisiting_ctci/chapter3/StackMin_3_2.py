"""
get the minimum value in a stack
"""

from stack import Stack


def MinStack(s):

	lowest = s.Data[0]
	
	for i in s.Data[1:]:
		if i < lowest:
			lowest = i
	return lowest


if __name__ == '__main__':

	test = Stack()
	for i in [9,3,1,4,7,5,5,]:
		test.Push(i)

	print(test)
	print('Minimum:\n')
	print(MinStack(test))		