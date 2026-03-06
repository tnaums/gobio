package eutils

import (
	"fmt"
	"net/http"
)

func (c *Client) EFetch(db, key, wenv string) (*http.Response, error) {
	url := fmt.Sprintf(baseURL+"efetch.fcgi?db=%s&query_key=%s&WebEnv=%s&rettype=fasta&retmode=text", db, key, wenv)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating req: %s", err)
		return &http.Response{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return &http.Response{}, err
	}
	return resp, nil
}
