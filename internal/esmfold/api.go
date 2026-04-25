package esmfold

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/tnaums/gobio/internal/protein"
)

func (c *Client) GetStructure(protein protein.Protein) (*http.Response, error) {
	fmt.Println(protein)

	url := baseURL
	var sequence = []byte(protein.AminoAcid)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sequence))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
