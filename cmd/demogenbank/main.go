package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	f, _ := os.Open("sequences/sequence.gb")
	g := dna.NewGenBank(f)

	for _, orf := range g.Sequence.Orfs {
		if len(orf.AminoAcid) > 150 {
			fmt.Println(orf)
		}
	}

}
