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

	// Get []PDWithPeptides where each entry is one protein
	// with Accession and identified peptides map
	pdSlice, _ := proteomediscoverer.ParseCSVWithPeptides(f)

	// Build string with all accession numbers for ncbi request
	builder := strings.Builder{}
	for _, record := range pdSlice {
		builder.WriteString(fmt.Sprintf("%s,", record.Accession))
	}
	
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost(builder.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body)


	for i := 0; i < len(pdSlice); i++ {
		// print summary for each
		pdSlice[i].PrintSummary(<-proteins)
	}

}
