package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/dna"
	"github.com/tnaums/gobio/internal/komagataella"
)

func main() {
	file, err := os.Open("sequences/pTAN309.fa")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	dnas := dna.DNAChannelFasta(file)
	dna, ok := <-dnas
	if !ok {
		fmt.Println("Protein channel is empty.")
	}

	protein := komagataella.GetMatureProtein(dna)
	
	fmt.Println(protein)
}
