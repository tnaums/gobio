// Demonstrates use of eutils.EPost for retrieving dna fasta
// sequences from NCBI. The response body is then sent to
// dna.ChannelFromFasta where the seqeunces are returned as
// protein.Protein type through a go channel.
package main

import (
	"fmt"
	//"io"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/protein"	
)

func main() {
	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second, eutils.WithGenbank())
	// generate *http.Response from ncbi query
	//resp, err := eutilsClient.EPost("AIZ65945.1,QIR83317.1,194680922,50978626,28558982,9507199,6678417,")	
	resp, err := eutilsClient.EPost("AIZ65945.1,QIR83317.1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	proteins := protein.ChannelFromGenbank(resp.Body)
	fmt.Println()
	for p := range proteins {
		fmt.Println(p)
	}


}
