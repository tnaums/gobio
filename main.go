package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/protein"
)

func main() {
	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second)
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost("P02845")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body)

	// Print first protein
	protein := <-proteins
	fmt.Println(protein)
	protein.CreateTrypticPeptides()
	for idx, peptide := range protein.Peptides {
		fmt.Printf("%d. %+v\n", idx, peptide)
	}

	

}
