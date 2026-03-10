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
	resp, err := eutilsClient.EPost("AIZ65945.1,QIR83317.1,194680922,50978626,28558982,9507199,6678417")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Create a new go channel
	proteins := make(chan protein.Protein)
	// start go routine with http response body as io.Reader and proteins channel
	go protein.ProteinPipeFasta(resp.Body, proteins)

	for p := range proteins { // iterate over proteins returned from go channel
		fmt.Println(p)
		fmt.Println()
	}
}
