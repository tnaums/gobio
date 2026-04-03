package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	resp, err := uniprotClient.GetAccession("A0A0A7LRQ7", "text/x-fasta")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body)

	// Print protein
	fmt.Println(<-proteins)
	fmt.Println()

	// Second example: generate *http.Response from uniprot query
	// for complete flatfile.  
	resp, err = uniprotClient.GetAccession("A0A0A7LRQ7", "text/x-flatfile")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read response.Body into []byte
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print the x-flatfile record
	fmt.Println(string(data))

	// Third example: generate *http.Response from uniprot query
	// to get full record in json format. Information is then unmarshalled
	// into a uniprot.UniprotRecord struct.
	//resp, err = uniprotClient.GetAccession("A0A0A7LRQ7", "application/json")
	resp, err = uniprotClient.GetAccession("Q8NID8", "application/json")	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read response.Body into []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create uniprot.UniprotRecord
	var record uniprot.UniprotRecord
	err = json.Unmarshal(data, &record)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// print features
	for _, feature := range record.Features {
		fmt.Printf("Type:\t%s\n", feature.Type)
		fmt.Printf("Category:\t%s\n", feature.Category)
		fmt.Printf("Description:\t%s\n", feature.Description)
		fmt.Printf("Begin:\t%s\n", feature.Begin)
		fmt.Printf("End:\t%s\n", feature.End)
		fmt.Printf("Molecule:\t %s\n", feature.Molecule)
		fmt.Printf("Evidences:\t%s\n", feature.Evidences)
		fmt.Println()
	}

	// create protein.Protein struct
	p := protein.NewProtein(record.Accession, record.Sequence.Sequence)
	fmt.Println(p)

}
