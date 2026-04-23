// Parses protein accession numbers and identified peptides from
// proteome discoverer summaries.
package proteomediscoverer

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/protein"
)

type Manager struct {
	Records []ProteomeDiscoverer
}

type ProteomeDiscoverer struct {
	Accession string
	Protein   protein.Protein
	Peptides  map[string]int
}

// Prints sequence in fasta format with only mapped peptides visible
func (p ProteomeDiscoverer) String() string {
	resultString := p.Protein.AminoAcid
	spaces := []byte{32}
	withSpaces := slices.Repeat(spaces, len(resultString))

	for key, _ := range p.Peptides {
		re, _ := regexp.Compile(key)
		bounds := re.FindStringIndex(resultString)
		for i := bounds[0]; i < bounds[1]; i++ {
			withSpaces[i] = resultString[i]
		}
	}

	builder := strings.Builder{}
	builder.WriteString(">mapped_peptides\n")
	for idx, base := range string(withSpaces) {
		if idx == 0 {
			builder.WriteRune(base)
			continue
		}
		if idx%60 == 0 {
			builder.WriteString("\n")
			builder.WriteRune(base)
			continue
		}
		builder.WriteRune(base)

	}
	return builder.String()
}

// Parses proteome discoverer summary that contains identified peptide
// information. Includes downloading of protein sequences from NCBI.
func ParseCSV(f io.Reader) (Manager, error) {
	start := true
	r := csv.NewReader(f)
	var manager Manager
	current := ProteomeDiscoverer{
		Accession: "",
		Peptides:  make(map[string]int),
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return manager, err
		}
		if record[0] == "Checked" {  // header line to discard
			continue
		}
		if record[0] == "FALSE" {  // indicates new protein or first protein
			if start {
				current.Accession = record[3]
				start = false
				continue
			}
			manager.Records = append(manager.Records, current)
			current.Peptides = make(map[string]int)
			current.Accession = record[3]
		}
		if record[1] == "FALSE" {  // Peptide information
			peptide := record[3][4 : len(record[3])-4]
			current.Peptides[peptide] += 1
		}
	}

	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second)

	// Build string with all accession numbers for ncbi request
	builder := strings.Builder{}
	for _, record := range manager.Records {
		builder.WriteString(fmt.Sprintf("%s,", record.Accession))
	}

	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost(builder.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response.Body (io.Reader)
	proteins := protein.ProteinChannelFasta(resp.Body)

	// Add proteins
	for i := 0; i < len(manager.Records); i++ {
		manager.Records[i].Protein = <-proteins
	}

	return manager, nil
}
