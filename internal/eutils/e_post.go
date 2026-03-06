package eutils

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func (c *Client) EPost() (*http.Response, error) {
	// Download protein records corresponding to a list of GI numbers.
	params := EPost{
		Database: "protein",
		IdList:   "194680922,50978626,28558982,9507199,6678417",
	}

	// Assemble the epost URL
	post := fmt.Sprintf("epost.fcgi?db=%s&id=%s", params.Database, params.IdList)
	url := baseURL + post

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return &http.Response{}, err
	}

	// Post the epost URL
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return &http.Response{}, err		
	}
	defer resp.Body.Close()
	
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return &http.Response{}, err
	}

	// Parse WebEnv and QueryKey
	rQK, _ := regexp.Compile("<QueryKey>(.+)</QueryKey>")
	rWE, _ := regexp.Compile("<WebEnv>(.+)</WebEnv>")
	
	queryKey := (rQK.FindStringSubmatch(string(b)))
	webEnv := (rWE.FindStringSubmatch(string(b)))

	// Assemble the efetch URL
	url = fmt.Sprintf(baseURL+"efetch.fcgi?db=%s&query_key=%s&WebEnv=%s&rettype=fasta&retmode=text", params.Database, queryKey[1], webEnv[1])


	req, err = http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating req: %s", err)
		return &http.Response{}, err
	}
	// Post the efetch URL
	resp, err = c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return &http.Response{}, err
	}
	
	return resp, nil
	
}
