from linkedlist import Node, LinkedList


def addNums(ll1, ll2):
	out_ll = LinkedList()

	ll1.SetFront()
	ll2.SetFront()

	remainder = 0
	total = ll1.Data + ll2.Data

	if total < 10 :
		out_ll.Add(total)
	else:
		remainder = int(total/10)
		out_ll.Add((total - (remainder*10)))

	while (ll1.Node.Next != None) and (ll2.Node.Next != None) :

		ll1.MoveForward()
		ll2.MoveForward()

		total = ll1.Data + ll2.Data + remainder
		remainder = 0

		if total < 10 :
			out_ll.Add(total)
		else:
			remainder = int(total/10)
			out_ll.Add((total - (remainder*10)))

	if ll1.Node.Next != None : 
		while ll1.Node.Next != None :
			ll1.MoveForward()

			total = ll1.Data + remainder
			remainder = 0

			if total < 10 :
				out_ll.Add(total)
			else:
				remainder = int(total/10)
				out_ll.Add((total - (remainder*10)))


	elif ll2.Node.Next != None:
		while ll2.Node.Next != None :
			ll2.MoveForward()
			total = ll2.Data + remainder
			remainder = 0

			if total < 10 :
				out_ll.Add(total)
			else:
				remainder = int(total/10)
				out_ll.Add((total - (remainder*10)))


	if remainder > 0:
		out_ll.Add(remainder)

	out_ll.SetFront()
	return out_ll


def toNum(ll):
	""" take a linked list storing a number and return it as an integer"""

	ll.SetBack()
	num  = str(ll.Data)

	while ll.Node.Prev != None:
		ll.MoveBackward()
		num += str(ll.Data)

	return int(num)


if __name__ == '__main__':

	n1 = 2781

	ll1 = LinkedList()

	for n in str(n1):
		ll1.FrontAdd(int(n))

	n2 = 92330

	ll2 = LinkedList()

	for n in str(n2):
		ll2.FrontAdd(int(n))

	total_ll = addNums(ll1, ll2)

	int_1 = toNum(ll1)
	int_2 = toNum(ll2)

	int_total = toNum(total_ll)

	print(f'{int_1} + {int_2} == {int_total}:\n')
	print(int_1+int_2 == int_total)