package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	f, _ := os.Open("sequences/sequence.gb")
	g := dna.NewGenBank(f)

	// Print GenBank.Sequence in fasta format
	fmt.Println(g.Sequence)

	// loop through orfs of g.Sequence and print anything
	// larger than 150 amino acids
	for _, orf := range g.Sequence.Orfs {
		if len(orf.AminoAcid) > 150 {
			fmt.Println(orf)
		}
	}

}
