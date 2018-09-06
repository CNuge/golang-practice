/*

you have two linked lists, each which stores the components of
a number in reverse order, digit by digit

sum the numbers

i.e.

671
is in the linked list as:

1 -> 7 -> 6

*/


package main

import (
	"./linkedlist"
	"fmt"
)

// take two numbers represented as linked lists and sum them to create a single linked list
// of digits
func SumDigits( d1, d2 linkedlist.Node) linkedlist.Node {
	d1 = d1.Front()
	d2 = d2.Front()


	out_dig := linkedlist.Node{}
	remainder := 0

	for ((d1.Next != nil) && (d2.Next != nil)) {
		new_dig := d1.Data + d2.Data + remainder
		fmt.Println(new_dig)
		remainder = 0

		if new_dig < 10 {
			
			out_dig.Add(new_dig)
		}else{
			remainder = int(new_dig/10)
			out_dig.Add((new_dig-(remainder*10)))
		}

		//advance to digit
		d1 = *d1.Next
		d2 = *d2.Next
	}

	if d1.Next == nil && d2.Next == nil {
		// sum the last two digits and remainder, 
		// carry any remainder to a separate node
		new_dig := d1.Data + d2.Data + remainder
		remainder = 0

		if new_dig < 10 {
			out_dig.Add(new_dig)
		}else{
			remainder = int(new_dig/10)
			out_dig.Add((new_dig-(remainder*10)))
			out_dig.Add(remainder)
		}

	}else if d1.Next != nil {
		// add the remaining d1 digits to the linked list, plus any remainder
		for d1.Next != nil {
			new_dig := d1.Data + remainder
			remainder = 0

			if new_dig < 10 {
				out_dig.Add(new_dig)
			}else{
				remainder = int(new_dig/10)
				out_dig.Add((new_dig-(remainder*10)))
			}

			//advance to digit
			d1 = *d1.Next
		}	



	}else if d2.Next != nil {
		// add the remaining d2 digits to the linked list, plus any remainder

		for d2.Next != nil {
			new_dig := d2.Data + remainder
			remainder = 0

			if new_dig < 10 {
				out_dig.Add(new_dig)
			}else{
				remainder = int(new_dig/10)
				out_dig.Add((new_dig-(remainder*10)))
			}

			//advance to digit
			d2 = *d2.Next
		}

	} 

	out_dig = *out_dig.Next
	out_dig.Prev = nil	
	return out_dig

}


func main(){

	// d1 9231 
	d1 := linkedlist.Node{Data: 9}

	d1.FrontAdd(2)
	d1.FrontAdd(3)
	d1.FrontAdd(1)

	// d2 2719 
	d2 := linkedlist.Node{Data: 2}

	d2.FrontAdd(7)
	d2.FrontAdd(1)
	d2.FrontAdd(9)

	output := SumDigits(d1, d2)

	fmt.Println(d1.Front())
	fmt.Println(d2.Front())
	fmt.Println(output)


}

