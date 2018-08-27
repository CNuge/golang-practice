/*

given a string check if it is a permutation of a palindrome

i.e.

rcarace - is a permutation of racecar, which is a palindrome

*/

package main


import(
	"fmt"
)

func IsPalindromePerm(s string ) bool{
	char_map := make(map[rune]int)

	for _ , i := range(s){
		char_map[i] += 1
	}

	odd_count := 0

	for _ , v := range char_map {
		if (v % 2) != 0 {
			odd_count += 1
		}
		if odd_count > 1{
			return false
		}
	}

	return true
}


func main(){

	s1 := "rcarace"
	fmt.Println(s1)
	fmt.Println(IsPalindromePerm(s1))

	s2 := "carcart"
	fmt.Println(s2)
	fmt.Println(IsPalindromePerm(s2))

	s3 := "carcarcar"
	fmt.Println(s3)
	fmt.Println(IsPalindromePerm(s3))

}

// take the string, split into characters and sort

// iterate through the characters


// if next character equal the current character, add to count
	
// if not, finalize previous count, is it odd?
	// if odd add 1 to odd count, if odd count > 1 return false
	
	// continue the iteration through the character sets

