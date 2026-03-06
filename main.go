// Module gobio provides tools for reading
// and analyzing DNA sequences from fasta files.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/protein"
)

func main() {
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
	}
	fileName := os.Args[1]
	selected := make([]protein.Protein, 0)
	
	// Create protein pipe from proteome fasta file
	proteins := make(chan protein.Protein)
	// Open file to create *os.File which implements io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	// Create go routine with opened fasta file and go channel
	go protein.ProteinPipeFasta(file, proteins)

	for p := range proteins { // iterate over proteins that are returned from go channel
		if p.Mass > 15 && p.Mass < 20 { 
			selected = append(selected, p)
			fmt.Println(p)
			fmt.Println()
		}

	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))
	fmt.Println("----------------------------------------\n")

	// Initialize client for api request
	eutilsClient := eutils.NewClient(50 * time.Second)
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Create a new go channel
	proteins = make(chan protein.Protein)
	// start go routine with http response body as io.Reader and proteins channel
	go protein.ProteinPipeFasta(resp.Body, proteins)

	for p := range proteins { // iterate over proteins returned from go channel
		fmt.Println(p)
		fmt.Println()
	}
}


	// Parse SignalP data from mycocosm
	// signalPMap := map[int]protein.SignalP{}
	// sigFile, err := os.Open("genomes/Fusve2/signalp.tab")
	// if er != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
