// Module gobio provides tools for reading
// and analyzing DNA sequences from fasta files.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/signalp"
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

	file2, err := os.Open("genomes/Fusve2/signalp.tab")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file2.Close()

	signalPMap, _ := signalp.NewSignalPMap(file2)
	//	fmt.Println(signalPMap)
	// Create go routine with opened fasta file and go channel
	go protein.ProteinPipeFasta(file, proteins)

	for p := range proteins { // iterate over proteins that are returned from go channel
		// get protein int from p.Header
		fields := strings.Split(p.Header, "|")
		number, _ := strconv.Atoi(fields[2])
		s, ok := signalPMap[number]
		if ok { // protein is in the signalp list
			mHeader := p.Header + "|secreted"
			mStart := s.NnCutPos
			mSequence := p.AminoAcid[mStart:]
			mature := protein.NewProtein(mHeader, mSequence)
			if mature.Mass > 16 && mature.Mass < 19 {
				selected = append(selected, p)
				fmt.Println(mature)
				fmt.Println()
			}

		}
	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))
	fmt.Println("----------------------------------------\n")

	// toAdd := signalp.SignalP{20, 5, 20, 0.99}
	// myMap := signalp.SignalPMap{}
	// myMap[5] = toAdd
	// fmt.Println(myMap)
	// // Initialize client for api request
	// eutilsClient := eutils.NewClient(50 * time.Second)
	// // generate *http.Response from ncbi query
	// resp, err := eutilsClient.EPost()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer resp.Body.Close()

	// // Create a new go channel
	// proteins = make(chan protein.Protein)
	// // start go routine with http response body as io.Reader and proteins channel
	// go protein.ProteinPipeFasta(resp.Body, proteins)

	// for p := range proteins { // iterate over proteins returned from go channel
	// 	fmt.Println(p)
	// 	fmt.Println()
	// }
}

// Parse SignalP data from mycocosm
// signalPMap := map[int]protein.SignalP{}
// sigFile, err := os.Open("genomes/Fusve2/signalp.tab")
// if er != nil {
// 	fmt.Fprintln(os.Stderr, err)
// 	os.Exit(1)
// }
