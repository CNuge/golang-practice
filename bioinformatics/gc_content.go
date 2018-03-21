package main

import (
	"fmt"
)

type dna_seq string

// take in a string, return the %gc content
func gc_content(dna_seq string) float64 {
	total_bp := len(dna_seq)
	gc := 0
	for _, char := range dna_seq {
		if char == 'G' || char == 'C' {
			gc++
		}
	}
	return float64(gc) / float64(total_bp) * 100.0
}

func main() {
	seq1 := "ATGAT"          // 1 gc .2
	seq2 := "ATGC"           // 2 gcs .5
	seq3 := "GCCTTTTTT"      //3 gcs .33
	seq4 := "GCGCGCGCGCGCGC" // 100%

	fmt.Println(gc_content(seq1))
	fmt.Println(gc_content(seq2))
	fmt.Println(gc_content(seq3))
	fmt.Println(gc_content(seq4))
}

// the commented out version does the same thing via a for loop,
// the range use above is a little more concise though
// also note the above one doesn't run if the _ is switched for a name

/*
func gc_content(dna_seq string) float64 {
	total_bp := len(dna_seq)
	gc := 0
	for i := 0; i < len(dna_seq); i++ {
		switch dna_seq[i] {
		case 'g':
			gc++
		case 'G':
			gc++
		case 'c':
			gc++
		case 'C':
			gc++
		default:
			continue
		}
	}
	return float64(gc)/float64(total_bp) * 100.0
}
*/
