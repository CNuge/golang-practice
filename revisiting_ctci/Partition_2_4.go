/*
take a linked list and partition it around a value

everything smaller than the value is left of all nodes greater than the value
the left and right don't need to be in order

i.e.

3 -> 5 -> 8 -> 5 -> 10 -> 1 -> 2
partition value = 5

1 -> 2 -> 3 -> 5 -> 8 -> 5 -> 10 


unclear:
whether there are instances where things would be mixed up like

 5 -> 3 -> 8 -> 5 

will work on the assumption this is not the case.
check when first answer is completed
*/