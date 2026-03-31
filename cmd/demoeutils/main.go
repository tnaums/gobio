// Demonstrates use of eutils.EPost for retrieving protein fasta
// sequences from NCBI. The response body is then sent to
// protein.ProteinChannelFasta where the seqeunces are returned as
// protein.Protein type through a go channel.
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
	resp, err := eutilsClient.EPost("AIZ65945.1,QIR83317.1,194680922,50978626,28558982,9507199,6678417,")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body) 

	// Print first protein
	fmt.Println(<-proteins)
	fmt.Println()

	// Print sequence from second protein
	p2 := <-proteins
	fmt.Println(p2.AminoAcid)
	fmt.Println()

	// For remaining proteins, print header, mass, sequence length
	for p := range proteins { 
		fmt.Printf(">%s|%.2fkDa|%daa", p.Header, p.Mass, len(p.AminoAcid))
		fmt.Println()
	}
}
