/*

read in a fasta file and store it in a new data structure
- this is the simple version that slurps up the entire file in memory and the parses
	the string to create the sequence objects

*/

package main

import(
	"fmt"
	"flag"
	"strings"
	"log" // for logging errors
	"io/ioutil" //input/output utilities https://golang.org/pkg/io/ioutil/ 
)


// represent a sequence
type sequence struct {
	name string
	seq string
}

// the function to return the sequence in fasta format when printed
func (sq sequence) String() string{
	return fmt.Sprintf(">%v\n%v\n", sq.name, sq.seq)
}

func parse_fasta(fasta_entry string) sequence {
	entry := strings.Split(fasta_entry, "\n")
	// first position is the name,
	// join everything but the first line into a single string
	return sequence{ name : entry[0],
					 seq : strings.Join(entry[1:], "")}
}


func main() {
	// this is the slice to hold all of the sequences 
	dna_sequences := []sequence{}
	// parse the command line arguments
	flag.Parse() 

	// iterate through all of the passed in files using flag.Args()
	for _, filename := range flag.Args() {
		fmt.Println(filename)
		//Opening a file to obtain an os.File value
		file, err := ioutil.ReadFile(filename)
		// check if that caused an error
    	if err != nil {
			log.Fatal(err)
		}

		// split the input file on the new seq characters
		data := strings.Split(string(file) , ">")

		// the first position is empty because of the leading >
		// so we iterate from 1:end
		for _ , entry := range data[1:] {
			fa_seq := parse_fasta(entry)
			dna_sequences = append(dna_sequences, fa_seq)
		}
	}

	// calling print on a sequence causes the String() method to be triggered
	// and it therefore pretty prints
	fmt.Println(dna_sequences[0])

}