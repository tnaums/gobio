package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/localblast"
	"github.com/tnaums/gobio/internal/uniprot"
)

func main() {
	accessions := []string{"Q876W5", "I1S3A5", "I1RPD9", "I1S5J8", "I1RHP3", "I1RR40", "I1RQV2"}
	// Initialize client for api request
	uniprotClient := uniprot.NewClient(15 * time.Second)

	for _, accession := range accessions {
		// Returns UniprotComplete which contains both unmarshaled info
		// from json and formatted x-flatfile for display
		record, err := uniprotClient.GetAccession(accession)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create fasta record
		p := record.GetFasta()
		blast := localblast.LocalBlast(p, "graminearum.ncbi.aa.fasta")

		fmt.Printf("Accession: %s\n", accession)
		fmt.Printf("%s\n\n", blast.BlastOutputIterations.Iteration.IterationHits.Hit[0].HitDef)
	}
}

// Q876W5 => FGSG_03544
// I1S3A5 => FGSG_11280
// I1RPD9 => FGSG_05906
// I1S5J8 => FGSG_12119
// I1RHP3 => FGSG_03304
// I1RR40 => FGSG_06549
// I1RQV2 => FGSG_06452
