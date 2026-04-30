package alphafold

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client method that retrieves json summary for a uniprot
// id. Unmarshals response into type AlphafoldSummary. Can be called
// directly by client program or used indirectly by call to client.GetCIF(id).
func (c *Client) GetSummaries(id string) (AlphafoldSummary, error) {
	fmt.Printf("Retrieving info for id: %s\n", id)

	url := baseURL + id

	var afSummary AlphafoldSummary

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AlphafoldSummary{}, fmt.Errorf("building request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AlphafoldSummary{}, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AlphafoldSummary{}, fmt.Errorf("reading body: %w", err)
	}

	err = json.Unmarshal(body, &afSummary)
	if err != nil {
		return AlphafoldSummary{}, fmt.Errorf("unmarshalling json: %w", err)
	}

	return afSummary, nil
}

// Client method to retrieve structure file for a uniprot id. There
// may be more than one structure for one id, maybe. Function returns only
// one but prints a warning message if there are more.
func (c *Client) GetCIF(id string) (*http.Response, error){
	summaries, err := c.GetSummaries(id)
	if err != nil {
		return nil, fmt.Errorf("getting summaries: %w", err)
	}

	if len(summaries) == 0 {
		return nil, fmt.Errorf("no structures found for %s", id)
	}
	if len(summaries) > 1 {
		fmt.Printf("There are %d structures for %s\n", len(summaries), id)
	}
	req, err := http.NewRequest("GET", summaries[0].CifURL, nil)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	return resp, nil
}
