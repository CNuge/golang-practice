/*
read in the an rna seqenece,
find the reference frame by searching for a start codon 
and transcribe it until a stop codon is reached
*/
package main

import(
	"fmt"
)

// the translation map RNA -> AA
var trans_map = map[string]string{"UUU":"F", "UUC":"F", "UUA":"L", "UUG":"L",
									"UCU":"S", "UCC":"S", "UCA":"S", "UCG":"S",
									"UAU":"Y", "UAC":"Y", "UAA":"STOP", "UAG":"STOP",
									"UGU":"C", "UGC":"C", "UGA":"STOP", "UGG":"W",
									"CUU":"L", "CUC":"L", "CUA":"L", "CUG":"L",
									"CCU":"P", "CCC":"P", "CCA":"P", "CCG":"P",
									"CAU":"H", "CAC":"H", "CAA":"Q", "CAG":"Q",
									"CGU":"R", "CGC":"R", "CGA":"R", "CGG":"R",
									"AUU":"I", "AUC":"I", "AUA":"I", "AUG":"M",
									"ACU":"T", "ACC":"T", "ACA":"T", "ACG":"T",
									"AAU":"N", "AAC":"N", "AAA":"K", "AAG":"K",
									"AGU":"S", "AGC":"S", "AGA":"R", "AGG":"R",
									"GUU":"V", "GUC":"V", "GUA":"V", "GUG":"V",
									"GCU":"A", "GCC":"A", "GCA":"A", "GCG":"A",
									"GAU":"D", "GAC":"D", "GAA":"E", "GAG":"E",
									"GGU":"G", "GGC":"G", "GGA":"G", "GGG":"G",}

// a type to represent the transcript in both formats
type transcript struct{
	rna string
	protein string
}

// find the First AA after the 
func (tr *transcript) Frame() int {
	// start scan of rna with [0:3] if not AUG, add one to index
	for index := 0 ; index < len(tr.rna); index++ {
		if tr.rna[index:index + 3] == "AUG" {
			start_pos := index + 3
			return start_pos
		}
	}
	return len(tr.rna)
}

// start with an initial bp scan to find "AUG"
func (tr *transcript) Translate() string {
	start := tr.Frame() // find the first AA's start position
	fmt.Printf("Start position: %v\n", start)
	for pos := start ; pos < len(tr.rna) -3 ; pos = pos + 3 {
		aa := trans_map[tr.rna[pos: pos+3]]
		if aa == "STOP" {
			break
		} else {
			tr.protein = fmt.Sprintf("%v%v", tr.protein, aa )
		}
	}
	return tr.protein
}



func main(){
	transcript1 := transcript{ rna: "AAAUAUGGCCGCAGAGUAGGGGGG" }
	transcript2 := transcript{ rna:  "AAUGUUUUUCCUAGCAAAAUGAAAA"}
	transcript1.Translate()
	fmt.Println(transcript1.protein)
	transcript2.Translate()
	fmt.Println(transcript2.protein)

//errors to address:
	// doesn't handle stop codons at the end of the string (this causes the error: tr.rna[pos: pos+3] )
	// need to handle overflow (no stop codon)
	// need to handle no start codon
	// need to have a Stringer() method for 

}