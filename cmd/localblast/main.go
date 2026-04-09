package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/localblast"
	"github.com/tnaums/gobio/internal/protein"	
)

func main() {
	// Open file with query sequence
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Open file to create *os.File
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a channel of Proteins
	proteins := protein.ProteinChannelFasta(file)
	protein, _ := <-proteins

	// local blast against Fusarium graminearum
	blast := localblast.LocalBlast(protein, "Ccarb.aa.fasta")
	// print results
	localblast.PrintBlastp(blast)

}
