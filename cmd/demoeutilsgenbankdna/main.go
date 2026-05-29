// Demonstrates use of eutils.EPost for retrieving dna fasta
// sequences from NCBI. The response body is then sent to
// dna.ChannelFromFasta where the seqeunces are returned as
// protein.Protein type through a go channel.
package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
)

func main() {
	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second, eutils.WithGenbank(), eutils.WithNucleotide())
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost("KM492932.1,MN339473.1") // dna sequences

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	fmt.Println(string(buf))

}
