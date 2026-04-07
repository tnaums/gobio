package main

import (
	"fmt"
	"os"
	"time"

	//"github.com/tnaums/gobio/internal/protein"
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
}
