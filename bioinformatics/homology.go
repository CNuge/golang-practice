package main



import(
	"fmt"
	"strings"
	)

//"os"

/*
By counting the number of differences between two homologous DNA 
strands taken from different genomes with a common ancestor, 
we get a measure of the minimum number of point mutations that 
could have occurred on the evolutionary path between the two strands.

This is called the 'Hamming distance'.

It is found by comparing two DNA strands and counting how many 
of the nucleotides are different from their equivalent in the other string.


goal for output:

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^
identity = x / y
%identity = %

*/

func homology_string(seq1, seq2 string) []string {
	hom_str := []string{}

	for i:= 0; i < len(seq1) ; i++ {
		if seq1[i] == seq2[i]{
			hom_str = append(hom_str, " ")
		} else {
			hom_str = append(hom_str, "^")			
		}
	}

	return hom_str
}


func homology_count(homo_string []string) int {
	matches := 0
	for _, v := range homo_string {
		if v == " " {
			matches++
		}
	}
	return matches
}

func main(){
	// could turn this into an input receiving program
	//seq1 := os.Args[1]
	//seq2 := os.Args[2]
	seq1 := "GAGCCTACTAACGGGAT"
	seq2 := "CATCGTAATGACGGCCT"


	homo_string := homology_string(seq1, seq2)

	homo_num := homology_count(homo_string) 

	fmt.Println(seq1)
	fmt.Println(seq2)
	fmt.Println(strings.Join(homo_string, ""))
	fmt.Printf("number of matches: %v / %v \n", homo_num, len(homo_string) )
	

	fmt.Printf("%.2f%%  identity\n", float64(homo_num)/float64(len(homo_string)) *100 )

}







