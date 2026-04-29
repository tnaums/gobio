package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/alphafold"
)

func main() {
	// Set the uniprot id for the structure
	id := "I1S3A5"

	// Initialize client for api request
	alphafoldClient := alphafold.NewClient(15 * time.Second)

	// Returns *http.Response with json summary
	resp, err := alphafoldClient.GetStructure(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}
