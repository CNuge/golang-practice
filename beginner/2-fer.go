package main

import(
	"fmt"
	"os"
)

/* this program takes and input name and makes it return the sentence
 as output,
 I am giving a default argument of 'you' which changes if there is a 
 name input to the command line following the program

	Note: this has led to an interesting discovery... which is obvious in hingsight
	the first os.Args (i.e. position 0) is ./2-fer

	this should have been obvious given that go is 0 indexed and I was geting the name using
	os.Args[1]

	Note2: if I use 'go run' as oppoesed to 'go build' then run, the first
	argument is very long, gives the path to where the on the fly compilation is occurring
*/
func main(){

	name := "you"

	// these revealed the details discussed in the notes
	//fmt.Println(len(os.Args))
	//fmt.Println(os.Args[0])

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Printf("One for %v, one for me.\n", name)

}