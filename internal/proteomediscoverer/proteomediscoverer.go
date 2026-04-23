// Parses protein accession numbers and identified peptides from
// proteome discoverer summaries.
package proteomediscoverer

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/tnaums/gobio/internal/protein"
)

type ProteomeDiscoverer struct {
	Accession string
}

type PDWithPeptides struct {
	Accession string
	Peptides  map[string]int
}

// Prints protein sequence and sequence with identified amino acids
// converted to 'x'.
func (p PDWithPeptides) PrintSummary(protein protein.Protein) {
	resultString := protein.AminoAcid
	resultBytes := []byte(protein.AminoAcid)
	fmt.Println(protein)
	for key, _ := range p.Peptides {
		re, _ := regexp.Compile(key)
		bounds := re.FindStringIndex(resultString)
		for i := bounds[0]; i < bounds[1]; i++ {
			resultBytes[i] = 120 // change to 'x'
		}
	}
	builder := strings.Builder{}
	builder.WriteString(">mapped_peptides\n")
	for idx, base := range string(resultBytes) {
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
	fmt.Println(builder.String())
}


func GetAccession(entry string) (ProteomeDiscoverer, error) {
	columns := strings.Split(entry, ",")
	fmt.Println(len(columns))
	return ProteomeDiscoverer{
		Accession: columns[3],
	}, nil
}

// Parses proteome discoverer summary that contains identified peptide
// information.
func ParseCSVWithPeptides(f io.Reader) ([]PDWithPeptides, error) {
	start := true
	r := csv.NewReader(f)
	var records []PDWithPeptides
	current := PDWithPeptides{
		Accession: "",
		Peptides:  make(map[string]int),
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return records, err
		}
		if record[0] == "Checked" {
			// header line to discard
			continue
		}
		if record[0] == "FALSE" {
			// indicates new protein or first protein
			if start {
				current.Accession = record[3]
				start = false
				continue
			}
			records = append(records, current)
			current.Peptides = make(map[string]int)
			current.Accession = record[3]
		}
		if record[1] == "FALSE" {
			peptide := record[3][4 : len(record[3])-4]
			current.Peptides[peptide] += 1
		}
	}
	return records, nil
}

// Parses simple summary that does not have peptide information
func ParseCSV(f io.Reader) ([]ProteomeDiscoverer, error) {
	r := csv.NewReader(f)
	var records []ProteomeDiscoverer
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return records, err
		}
		if record[0] == "Checked" {
			continue
		}

		pd := ProteomeDiscoverer{
			Accession: record[3],
		}
		records = append(records, pd)
	}
	return records, nil
}
