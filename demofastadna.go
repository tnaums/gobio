// This script demonstrates how to work with dna sequences
// from fasta files containing multiple entries.
package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	// Create a channel for sending Dna
	dnach := make(chan dna.Dna)

	fmt.Println("Welcome to gobio!")

	// Get filename from command line
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Open file to create *os.File which implements io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Create go routine with opened fasta file and go channel
	go dna.DnaPipeFasta(file, dnach)

	// Retrieve first sequence and print
	first :=  <- dnach
	fmt.Println(first)

	// Retrieve second sequence and print
	second := <- dnach
	fmt.Println(second)

	// Iterate over orfs and print if over 100 amino acids
	for idx, orf := range second.Orfs {
		if len(orf.AminoAcid) > 100 {
			fmt.Println(idx)
			fmt.Println(orf)
			fmt.Println()
		}
	}

	// Iterate over remaining dna sequences and print the
	// header line and sequence length
	for d := range dnach {
		fmt.Println(d.Header)
		fmt.Printf("Length: %d\n", len(d.Parent))
	}
}
