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

	data := FastaParser(file)
	for i := 0; i < len(data); i = i + 2 {
		newProtein := Protein{
			Header:    data[i],
			AminoAcid: data[i+1],
		}
		returnSlice = append(returnSlice, newProtein)
	}
	return returnSlice, nil
}

// The FastaParser function reads a fasta file, extracts
// the sequence name from the header, and creates a sequence
// string from the sequence. Returns a slice of strings with
// header followed by associated sequence.
func FastaParser(r io.Reader) (data []string) {
	start := true
	name := ""
	sequence := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ">"){
			if !start {
				data = append(data, sequence)
				sequence = ""
			}
			name = scanner.Text()
			data = append(data, name[1:])
			start = false
		} else {
			sequence += scanner.Text()
		}
	}
	data = append(data, sequence)
	return data

}

// ProteinPipeFasta reads fasta sequences from an io.Reader interface,
// such as an *os.File returned from os.Open(fileName).
// Returns stream of Protein structs through the provided go channel.
// Once the last Protein is sent, closes the channel.
func ProteinPipeFasta(r io.Reader, out chan<- Protein) {
	start := true
	name := ""
	sequence := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ">"){
			if !start {
				out <- Protein{name, sequence}
				sequence = ""
			}
			name = scanner.Text()
			name = name[1:]
			start = false
		} else {
			sequence += scanner.Text()
		}
	}
	out <- Protein{name, sequence}
	close(out)
}
