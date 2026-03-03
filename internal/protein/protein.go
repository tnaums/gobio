// Package protein provides a protein type to store protein
// sequence information
package protein

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// This initial version of the Protein struct contains only
// the protein sequence and the fasta header
type Protein struct {
	Header    string
	AminoAcid string
}

// NewProteinFromFasta is a function that creates a
// slice of type Protein from a fasta file that contains
// one or more protein sequences
func NewProteinFromFasta(filename string) ([]Protein, error) {
	returnSlice := make([]Protein, 0)
	file, err := os.Open(filename)
	if err != nil {
		return []Protein{}, err
	}

	header, sequence := FastaParser(file)
	newProtein := Protein{
		Header:    header,
		AminoAcid: sequence,
	}
	returnSlice = append(returnSlice, newProtein)

	return returnSlice, nil
}

// The FastaParser function reads a fasta file, extracts
// the sequence name from the header, and creates a sequence
// string from the sequence.
func FastaParser(r io.Reader) (name, sequence string) {
	name = ""
	sequence = ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ">"){
			name = scanner.Text()
			name = name[1:]
		} else {
			sequence += scanner.Text()
		}
	}
	return name, sequence

}
