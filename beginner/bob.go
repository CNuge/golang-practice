package main

import(
	"fmt"
	"os"
)


func Bob(input string) string {
	if len(input) == 0 {
		return "Fine. Be that way!"
	} else if input[len(input)-1] == '!' {
		return "Whoa, chill out!"
	} else if input[len(input)-1] == '?' {
		if input[len(input)-2] != '!' {
			return "sure"
		} else {
			return "Calm down, I know what I'm doing!"
		}
	} else {
		return "Whatever"
	}
}


func main(){
	fmt.Println(os.Args[1])
	output := Bob(os.Args[1])

	fmt.Println(output)
}