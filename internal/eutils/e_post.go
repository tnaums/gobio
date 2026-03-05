package eutils

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func (c *Client) EPost() (string, string, error) {

	params := EPost{
		Database: "protein",
		IdList:   "194680922,50978626,28558982,9507199,6678417",
	}

	post := fmt.Sprintf("epost.fcgi?db=%s&id=%s", params.Database, params.IdList)
	url := baseURL + post

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return "", "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return "", "", err
	}
	defer resp.Body.Close()
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	// Parse WebEnv and QueryKey
	rQK, _ := regexp.Compile("<QueryKey>(.+)</QueryKey>")
	rWE, _ := regexp.Compile("<WebEnv>(.+)</WebEnv>")
	
	queryKey := (rQK.FindStringSubmatch(string(bodyBytes)))
	webEnv := (rWE.FindStringSubmatch(string(bodyBytes)))
	return queryKey[1], webEnv[1], nil
}
