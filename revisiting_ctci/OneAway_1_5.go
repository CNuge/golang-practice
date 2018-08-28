/*

given two strings, write a method to determine if they are one (or less)
change (insertion, deletion, replacement) away from one another

*/


package main

import (
	"fmt"
)


// take two strings and compare them to one another to see if they are one
// insertion, deletion or replacement away from one another
func OneAway(s1, s2 string) bool {
	diff := 0
	long := ""
	short := ""
	
	// check for replacement first, only valid if same length
	if len(s1) == len(s2){
		for i:=0 ; i < len(s1) ; i++ {
			if s1[i] != s2[i]{
				diff++
			}
			if diff > 1 {
				return false
			}
		}

	// if they aren't the same size only one indel is allowed
	// check to see if they are one different in length, determine the larger
	} else if (len(s1) - len(s2)) == 1 {
		long = s1
		short = s2
	} else if (len(s1) - len(s2)) == -1 {
		long = s2
		short = s1
	} else {
		return false
	}

	offset := 0
	for i := 0  ; i < len(short); i++ {
		li := i + offset
		if short[i] != long[li] {
			offset++
			diff++
		}
		if diff > 1 {
			return false
		}
	}

	return true
}




func main() {

	t1 := "TATTA"
	t2 := "TACTTA"
	fmt.Printf("%v\t%v\t%v\n", t1, t2, OneAway(t1,t2))

	t3 := "CGCGCG"
	t4 := "CGCTCG"
	fmt.Printf("%v\t%v\t%v\n", t3, t4, OneAway(t3,t4))


	t5 := "ATGCATGC"
	t6 := "ATGCTTATGC"
	fmt.Printf("%v\t%v\t%v\n", t5, t6, OneAway(t5,t6))


	t7 := "GGGGTAC"
	t8 := "GGCCTAC"
	fmt.Printf("%v\t%v\t%v\n", t7, t8, OneAway(t7,t8))


}
