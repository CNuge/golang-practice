
/*
take a string and compress it into repeated digit counts:

AAAATTTTTGGG
A4T5G3

if the new string is longer than the original, then return the original.

*/

package main

import (
	"fmt"
)

func CompressString(s string) string {
	outstring := ""

	var current_letter rune
	var current_count int

	for _, char := range s {
		if char == current_letter{
			current_count += 1
		}else{
			if current_count != 0 {
				addition := fmt.Sprintf("%v%v", string(current_letter), current_count)
				outstring += addition
			}
			current_letter = char
			current_count = 1

		}
	}
	addition := fmt.Sprintf("%v%v", string(current_letter), current_count)
	outstring += addition

	if len(outstring) < len(s){
		return outstring
	}
	return s
}


func main() {

	t1 := "AAAATTTTTGGG"
	o1 := "A4T5G3"

	fmt.Println(t1)
	fmt.Println(CompressString(t1))
	fmt.Println(o1)
	fmt.Println(CompressString(t1) == o1)

	t2 := "ATGCATGCATGC"
	o2 := "ATGCATGCATGC"

	fmt.Println(t2)
	fmt.Println(o2)
	fmt.Println(CompressString(t2) == o2)

	t3 := "CGGGAAAATTTTTGGG"
	o3 := "C1G3A4T5G3"

	fmt.Println(t3)
	fmt.Println(CompressString(t3))
	fmt.Println(o3)
	fmt.Println(CompressString(t3) == o3)

}