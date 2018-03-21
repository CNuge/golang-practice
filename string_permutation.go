/*
given two strings, see if one is a permuation of the other
*/

package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
 how to sort an array in standard library:
	x := []string{"B","C","A"}
	sort.Strings(x)
	fmt.Println(x)
sort:
https://golang.org/pkg/sort/
*/

// sort and test equality
func StrEq(str1, str2 string) bool {
	s_str1 := strings.Split(str1, "")
	s_str2 := strings.Split(str2, "")

	sort.Strings(s_str1)
	sort.Strings(s_str2)

	return reflect.DeepEqual(s_str1, s_str2)
}

func main() {
	str1 := "CAM NUGE"
	str2 := "MAC NUGE"
	str3 := "GREG NUGE"

	fmt.Println(StrEq(str1, str2))
	fmt.Println(StrEq(str1, str3))

}
