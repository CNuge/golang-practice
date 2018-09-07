from linkedlist import LinkedList, Node


def KthToLast(ll, n):
	ll.SetBack()

	i = 1
	while i < n:
		ll.MoveBackward()
		i += 1

	return ll





if __name__ == '__main__':

	n1 = 987654321

	ll1 = LinkedList()

	for n in str(n1):
		ll1.Add(int(n))

	print(ll1)
	x = KthToLast(ll1, 3)
	print(x.Data)

