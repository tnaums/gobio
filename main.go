// Module gobio provides tools for reading
// and analyzing DNA sequences from fasta files.
package main

import (
	"fmt"
	"os"
	//	"github.com/tnaums/gobio/internal/dna"
	"github.com/tnaums/gobio/internal/protein"	
)

func main() {
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
	}

	fileName := os.Args[1]
	// dnaStruct, err := dna.NewDnaFromFasta(fileName)
	// if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// }

	// for _, orf := range dnaStruct.Orfs {
	// 	if len(orf.AminoAcid) > 200 {
	// 		fmt.Println(orf)
	// 	}
	// }

	// p, err := protein.NewProteinFromFasta("sequences/test_file.fa")
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// for _, protein := range p {
	// 	fmt.Println(protein.Header)
	// 	fmt.Println(protein.AminoAcid)
	// 	fmt.Println()
	// }

	// Parse SignalP data from mycocosm
	// signalPMap := map[int]protein.SignalP{}
	// sigFile, err := os.Open("genomes/Fusve2/signalp.tab")
	// if er != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	

	
	// Create protein pipe from proteome fasta file
	proteins := make(chan protein.Protein)
	selected := make([]protein.Protein, 0)
	
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	go protein.ProteinPipeFasta(file, proteins)

	for p := range proteins {
		if p.Mass > 15 && p.Mass < 20 {
			selected = append(selected, p)
			fmt.Println(p.Header)
			fmt.Println(p.AminoAcid)
			fmt.Println(p.Mass)
			fmt.Println()
		}

	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))

}
