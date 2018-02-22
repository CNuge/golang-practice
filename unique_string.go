/*
write an algorithm to determine is a string has all unique characters
*/

package main

import (
	"fmt"
	)

// make a map and check if next character is in the map
func IsUnique(str string) bool {
	seen := make(map[rune]bool)
	for _, i := range str {
		if seen[i] == true {
			return false
		} else {
			seen[i] = true
		}
	}
	return true
}




func main(){
	str1 := "CAMERON"
	str2 := "SAMANTHA"

	fmt.Println(IsUnique(str1))
	fmt.Println(IsUnique(str2))

}
