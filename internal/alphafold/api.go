package alphafold

import (
	"fmt"
	"net/http"
)

func (c *Client) GetStructure(id string) (*http.Response, error) {
	fmt.Printf("Retrieving info for id: %s\n", id)

	url := baseURL + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
