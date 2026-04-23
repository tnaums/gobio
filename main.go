package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/proteomediscoverer"
)

func main() {

	// Open Proteome Discoverer results summary
	f, err := os.Open("proteomediscoverer/money.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Parse the file. The returned Manager contains Records, a
	// slice of ProteomeDiscoverer type
	manager, _ := proteomediscoverer.ParseCSV(f)

	for i := 0; i < len(manager.Records); i++ {
		// First, print the protein in fasta format
		fmt.Println(manager.Records[i].Protein)
		// Second, print with mapped peptides only
		fmt.Println(manager.Records[i])
		fmt.Println("------------------------------------------------------------")
	}

}
