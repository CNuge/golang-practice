/*
go from DNA to RNA
ATGC -> UACG
*/

package main

import (
	"fmt"
	"strings"
)

type dna_seq string

func translation(dna dna_seq) string {
	rna := []string{}
	for _, bp := range dna {
		switch bp {
		case 'A':
			rna = append(rna, "U")
		case 'T':
			rna = append(rna, "A")
		case 'G':
			rna = append(rna, "C")
		case 'C':
			rna = append(rna, "G")
		}
	}
	return strings.Join(rna, "")
}

func main() {
	var seq1 dna_seq
	seq1 = "ATGCATGC"
	fmt.Println(seq1)
	fmt.Println(translation(seq1))
}
