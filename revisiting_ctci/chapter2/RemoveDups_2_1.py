from linkedlist import Node, LinkedList

def RemoveDups(ll):
	ll.SetFront()
	
	new_ll = LinkedList()
	new_ll.Add(ll.Node.Data)
	seen = [ll.Node.Data]
 
 	while ll.Node.Next != None:
 		ll.MoveForward()
 		if ll.Data not in seen:
 			new_ll.Add(ll.Data)
 			seen.append(ll.Data)

 	return new_ll


if __name__ == "__main__":

	n1 = 278182219652

	ll1 = LinkedList()

	for n in str(n1):
		ll1.Add(int(n))

	ll1

	ll1 = RemoveDups(ll1)
	
	ll1