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

// stringer to print the RNA and AA side by side
func (tr transcript) String() string{
	return fmt.Sprintf("rna:%v\nprotein:%v\n", tr.rna, tr.protein)
}

// find the First AA after the 
func (tr *transcript) Frame() int {
	// start scan of rna with [0:3] if not AUG, add one to index
	for index := 0 ; index <= len(tr.rna) - 3 ; index++ {
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
	if start == len(tr.rna){
		tr.protein = "..."
		return tr.protein
	}
	stop_found := 0
	for pos := start ; pos <= len(tr.rna) - 3 ; pos = pos + 3 {
		aa := trans_map[tr.rna[pos: pos+3]]
		if aa == "STOP" {
			stop_found = 1
			break
		} else {
			tr.protein = fmt.Sprintf("%v%v", tr.protein, aa )
		}
	}
	if stop_found == 0 {
		tr.protein = fmt.Sprintf("%v%v", tr.protein, "...")
	}
	return tr.protein
}


func main(){
	transcript1 := transcript{ rna: "AAAUAUGGCCGCAGAGUAGGGGGG" } // AAE
	transcript2 := transcript{ rna:  "AAUGUUUUUUUUUUUUUUUUUUUU"} // FFFFF...
	transcript3 := transcript{ rna:  "AAUGUUUAAAUUUAAAGAAUAG" } // FKFKE
	transcript4 := transcript{ rna:  "UUAAAUUUAAAGAAUAG" } // no start

	transcript1.Translate()
	fmt.Println(transcript1.protein)
	transcript2.Translate()
	fmt.Println(transcript2.protein)
	transcript3.Translate()
	fmt.Println(transcript3.protein)
	transcript4.Translate()
	fmt.Println(transcript4.protein)
	fmt.Println("\n\n\n\n")
	fmt.Println(transcript1)
//errors to address:
	// doesn't handle stop codons at the end of the string (this causes the error: tr.rna[pos: pos+3] )
	// need to handle overflow (no stop codon) - done
	// need to handle no start codon - done
	// need to have a Stringer() method for  -done

}