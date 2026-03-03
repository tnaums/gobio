// Module gobio provides tools for reading
// and analyzing DNA sequences from fasta files.
package main

import (
	"fmt"
	"os"
	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
	}

	fileName := os.Args[1]
	dnaStruct, err := dna.NewDnaFromFasta(fileName)
	if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
	}

	for _, orf := range dnaStruct.Orfs {
		if len(orf.AminoAcid) > 200 {
			fmt.Println(orf)
		}
	}

}
