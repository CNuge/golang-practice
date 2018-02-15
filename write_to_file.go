package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// check for file writing errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1.txt", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("dat2.txt")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	// Defer is used to ensure that a function call is performed later 
	// in a program’s execution, usually for purposes of cleanup.
	defer f.Close()

	// A `WriteString` is also available.
	f.WriteString("writes\n")

	fmt.Fprintf(os.Stdout, "this is written to standard output\n so we can catch it with an > or >>\n")


}