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

// represent a single sequence
type seq struct {
	name string
	sequence string
}

// represent a list of sequences
type Fasta struct {
	entries []seq
}

// add a seq struct instance to the fasta struct
func (sl *Fasta) AddItem(item seq) []seq {
	sl.entries = append(sl.entries, item)
	return sl.entries
}

// the function to return the sequence in fasta format when printed
func (sq seq) String() string {
	return fmt.Sprintf(">%v\n%v\n", sq.name, sq.sequence)
}

func ParseFasta(fasta_entry string) seq {
	entry := strings.Split(fasta_entry, "\n")
	// first position is the name,
	// join everything but the first line into a single string
	return seq{ name : entry[0],
				sequence : strings.Join(entry[1:], "")}
}

func ReadFasta(filename string) Fasta {
	fileseqs := Fasta{} // start an empty Fasta instance

	//Opening a file to obtain an os.File value
	file, err := ioutil.ReadFile(filename)
	// check if that caused an error
	if err != nil {
		log.Fatal(err)
	}

	// split the input file on the new seq characters
	data := strings.Split(string(file) , ">")

	// the first position is empty because of the leading >
	// so we iterate from 1:end and get the sequence
	// here we parse the fasta and add it to the slice of seq
	for _ , entry := range data[1:] {
		fileseqs.AddItem(ParseFasta(entry))
	}
	return fileseqs
}

func main() {
	// this is the slice to hold all of the entries 
	allfasta := Fasta{}
	// parse the command line arguments
	flag.Parse() 

	// iterate through all of the passed in files using flag.Args()
	for _, filename := range flag.Args() {
		fasta_seq := ReadFasta(filename)
		allfasta.entries = append(allfasta.entries, fasta_seq.entries...)
		// see https://golang.org/ref/spec#Appending_and_copying_slices 
		// for ... explination
	}

	// calling print on a sequence causes the String() method to be triggered
	// and it therefore pretty prints
	// show first
	fmt.Println(allfasta.entries[0])
	// show last
	fmt.Println(allfasta.entries[len(allfasta.entries)-1])

}

