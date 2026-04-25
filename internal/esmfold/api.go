package esmfold

import (
	"bytes"
	"fmt"
	"net/http"
)

func (c *Client) GetStructure(sequence string) (*http.Response, error) {
	fmt.Println(sequence)

	url := baseURL
	var sequenceStr = []byte(sequence)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sequenceStr))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
