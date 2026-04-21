package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/proteomediscoverer"
)

func main() {
	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second)

	// Open Proteome Discoverer results summary
	f, err := os.Open("proteomediscoverer/egg_glycoprotein.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get []PDWithPeptides slice where each entry is one protein
	// with Accession and identified peptides map
	p, _ := proteomediscoverer.ParseCSVWithPeptides(f)

	// generate *http.Response from ncbi query for first protein
	resp, err := eutilsClient.EPost(p[0].Accession)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body)

	// Get the protein from channel
	protein := <-proteins

	// Add Peptides attribute to protein
	protein.CreateTrypticPeptides()
	// iterate over peptides and print
	builder := strings.Builder{}
	for _, peptide := range protein.Peptides {
		if p[0].Peptides[peptide.Sequence] > 0 {
			builder.WriteString(peptide.Sequence)
			continue
		}
		builder.WriteString(strings.ToLower(peptide.Sequence))
	}
	fmt.Println(builder.String())

}
