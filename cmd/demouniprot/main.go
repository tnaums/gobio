package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/localblast"
	"github.com/tnaums/gobio/internal/uniprot"
)

func main() {
	// Initialize client for api request
	uniprotClient := uniprot.NewClient(15 * time.Second)

	// Returns UniprotComplete which contains both unmarshaled info
	// from json and formatted x-flatfile for display
	record, err := uniprotClient.GetAccession("A0A0A7LRQ7")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print features from record
	record.PrintFeatures()

	// Print fasta record
	fmt.Println(record.GetFasta())

	// Print complete flatfile
	fmt.Println(string(record.GetFlatFile()))

	// Create a list of accession numbers and retrieve
	accessions := []string{"Q8NID8", "Q876W5", "I1S3A5", "I1RPD9"}
	for _, accession := range accessions {
		fmt.Println()
		record, err := uniprotClient.GetAccession(accession)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create protein.Protein
		p := record.GetFasta()
		// Print fasta sequence
		fmt.Println(p)
		// Print features from record
		record.PrintFeatures()

		// perform local blast
		blast := localblast.LocalBlast(p, "graminearum.ncbi.aa.fasta")

		// print results
		localblast.PrintBlastp(blast)
	}
}
