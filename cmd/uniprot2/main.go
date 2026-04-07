package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/uniprot"
)

func main() {
	// Initialize client for api request
	uniprotClient := uniprot.NewClient(15 * time.Second)
	// First example: Generate *http.Response from uniprot query for fasta file
	// and create protein.Protein type
	buf, err := uniprotClient.GetAccessions([]string{"A0A0A7LRQ7", "Q8NID8"}, "text/x-fasta")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(buf)

	// Print protein
	for protein := range proteins{
		fmt.Println(protein)
		fmt.Println()
	}

}
