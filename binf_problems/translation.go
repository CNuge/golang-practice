/*
read in the an rna seqenece,
find the reference frame by searching for a start codon 
and transcribe it until a stop codon is reached
*/
package main

import(
	"fmt"
	"string"
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

// start with an initial bp scan to find "AUG"
func (tr *transcript) Translate(){
	index := 0 // start scan of rna with [0:4] if not AUG, add one to index
	
}



func main(){

}