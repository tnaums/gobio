// Demonstrates use of protein.ProteinChannelFasta for
// reading proteins from a fasta file containing one
// or more sequences.
package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/protein"
)

func main() {
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
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

	// Create a channel of Proteins
	proteins := protein.ProteinChannelFasta(file)

	protein, ok := <-proteins
	if ok {
		fmt.Println(protein)
	} else {
		fmt.Println("Protein channel is empty.")
	}

	protein, ok = <-proteins
	if ok {
		fmt.Println(protein)
	} else {
		fmt.Println("Protein channel is empty.")
	}	


	var count int
	for protein := range proteins {
		if protein.Mass > 200 {
			fmt.Println(protein.Header)
			fmt.Println()
			count++
		}
	}
	fmt.Printf("Found %d proteins larger than 20 kDa\n", count)
	fmt.Println("----------------------------------------\n")

}
