package uniprot

import (
	"fmt"
	"net/http"
)

func (c *Client) GetAccession(accession, responseType string) (*http.Response, error) {
	// Download fasta format sequences from list of accession numbers
	params := UniprotAPI{
		ResponseType: responseType,
		IdList:       accession,
	}

	// Assemble the URL
	url := baseURL + accession

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return &http.Response{}, err
	}

	req.Header.Set("Accept", params.ResponseType)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return &http.Response{}, err
	}
	
	return resp, nil
}
