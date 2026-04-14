package uniprot

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// For given accession, retrieves both a complete record as json
// that is unmarshalled into a go struct and complete record as
// flatfile that is saved for printing.
func (c *Client) GetAccession(accession string) (*UniprotComplete, error) {
	// Assemble the URL
	url := baseURL + accession

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return nil, err
	}

	// request data in json format
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making json request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// read json response into []byte
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading json response body: %v", err)

	}

	// unmarshal into uniprot.UniprotRecord
	var record UniprotRecord
	err = json.Unmarshal(data, &record)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json response: %v", err)
	}

	// request data in text/x-flatfile format
	req.Header.Set("Accept", "text/x-flatfile")
	resp, err = c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making flatfile request: %v", err)
	}
	defer resp.Body.Close()

	// read response.Body into []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading flatfile response: %v", err)
	}

	// assemble UniprotComplete from the two parts
	return &UniprotComplete{
		JSON:     record,
		Flatfile: data,
	}, nil
}

// func (c *Client) GetAccession(accession, responseType string) (*http.Response, error) {
// 	// Download sequence in requested format for one uniprot accession
// 	params := UniprotAPI{
// 		ResponseType: responseType,
// 	}

// 	// Assemble the URL
// 	url := baseURL + accession

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("Error creating request: %s\n", err)
// 		return &http.Response{}, err
// 	}

// 	req.Header.Set("Accept", params.ResponseType)
// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error making request: %s\n", err)
// 		return &http.Response{}, err
// 	}

// 	return resp, nil
// }

// func (c *Client) GetAccessions(accessions []string, responseType string) (*bytes.Buffer, error) {
// 	// Download sequences in requested format for list of accession numbers
// 	params := UniprotAPI{
// 		ResponseType: responseType,
// 	}

// 	// Create an empty bytes.Buffer
// 	var buf bytes.Buffer
// 	for _, accession := range accessions {
// 		// Assemble the URL
// 		url := baseURL + accession

// 		req, err := http.NewRequest("GET", url, nil)
// 		if err != nil {
// 			fmt.Printf("Error creating request: %s\n", err)
// 			return nil, err
// 		}

// 		req.Header.Set("Accept", params.ResponseType)
// 		resp, err := c.httpClient.Do(req)
// 		if err != nil {
// 			fmt.Printf("Error making request: %s\n", err)
// 			return nil, err
// 		}
// 		data, _ := io.ReadAll(resp.Body)
// 		buf.Write(data)
// 	}

// 	return &buf, nil
// }
