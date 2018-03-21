/*

get the reverse compliment of a DNA sequence,

print it out 5'->3'

*/

package main

import (
	"fmt"
	"strings"
)

var comp_map = map[string]string{
	"A": "T",
	"T": "A",
	"G": "C",
	"C": "G"}

// take a string of DNA as input, split it to an array and then
// iterate over the slice in reverse order, getting the complimentary bp
// from the map

func rev_comp(dna_string string) string {
	original := strings.Split(dna_string, "")
	rev_comp := []string{}
	for i := len(original) - 1; i >= 0; i-- {
		rev_comp = append(rev_comp, comp_map[original[i]])
	}
	return strings.Join(rev_comp, "")
}

func main() {
	test1 := "AAAAACGCGTATATAT"
	test2 := "GCGCGCGCGCGCGCGCGCAAATATATATGGCGGTTTTTTA"

	fmt.Println(test1)
	fmt.Println(rev_comp(test1))

	fmt.Println(test2)
	fmt.Println(rev_comp(test2))

	fmt.Println(comp_map["A"])

}
