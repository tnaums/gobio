package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/localblast"
	"github.com/tnaums/gobio/internal/proteomediscoverer"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <file.csv>")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Open file to create *os.File which is an io.ReadCloser
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Get slice of records
	manager, _ := proteomediscoverer.ParseCSV(file)

	// run localblast for selected protein against 3 databases
	proteomes := []string{"verticillioides.aa.fasta", "graminearum.aa.fasta", "subglutinans.aa.fasta", "proliferatum.aa.fasta", "Vdahliae.aa.fasta", "Cgram.fasta", "Ccarb.aa.fasta"}
	for idx, proteome := range proteomes {
		fmt.Printf(" %.2d. Performing blastp against %s\n", idx, proteome)
		blast := localblast.LocalBlast(manager.Records[4].Protein, proteome)
		localblast.PrintBlastp(blast)
	}
}
