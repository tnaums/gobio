// Demonstrates use of eutils.EPost for retrieving dna fasta
// sequences from NCBI. The response body is then sent to
// dna.ChannelFromFasta where the seqeunces are returned as
// protein.Protein type through a go channel.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second, eutils.WithNucleotide())
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost("KM492932.1,MN339473.1") // a dna sequence

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	dnas := dna.ChannelFromFasta(resp.Body) 

	// Print first dna
	fmt.Println(<-dnas)
	fmt.Println()
	// Print second dna
	fmt.Println(<-dnas)
	fmt.Println()

}
