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
	record, err := uniprotClient.GetAccession("Q876W5")
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


// Q876W5 => FGSG_03544
// I1S3A5 => FGSG_11280
// I1RPD9 => FGSG_05906
// I1S5J8 => FGSG_12119
// I1RHP3 => FGSG_03304
// I1RR40 => FGSG_06549
// I1RQV2 => FGSG_06452
