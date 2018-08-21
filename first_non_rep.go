/*
find the first non repeating character in a string
*/

package main


import(
	"fmt"
)


func FirstUniq(s string) string{
	chars := make(map[rune]int)
	keys := []rune{}
	for _, i := range s{
		keys = append(keys, i)
		chars[i] += 1

	}

	for _, k := range keys{
		if chars[k] == 1{
			return string(k)
		}
	}

	return ""
}




func main(){

	input := "this is the string we query"

	fmt.Println(FirstUniq(input))

}
