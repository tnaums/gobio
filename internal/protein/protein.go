// Package protein provides a protein type to store protein
// sequence information
package protein

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Contains header and amino acid sequence, parsed from
// fasta file. Mass can be calculated from AminoAcid by
// calling calculateMass(aaSequence)
type Protein struct {
	Header    string
	AminoAcid string
	Mass      float64
}

// Holds output of SignalP 6 analysis which predicts
// probability of protein secretion
type SignalP struct {
	NnCutPos  int
	NnVote    int
	HmmCutPos int
	HmmProb   float64
}

// NewProteinFromFasta creates a slice of type Protein from a fasta file
// containing one or more protein sequences.
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
			Mass:      calculateMass(data[i+1]),
		}
		returnSlice = append(returnSlice, newProtein)
	}
	return returnSlice, nil
}

// FastaParser reads a fasta file, extracts the sequence name from
// the header and creates a sequence string from the sequence.
// Returns a slice of strings with alternating header and sequence.
func FastaParser(r io.Reader) (data []string) {
	start := true
	name := ""
	sequence := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ">") {
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
		if strings.HasPrefix(scanner.Text(), ">") {
			if !start {
				out <- Protein{name, sequence, calculateMass(sequence)}
				sequence = ""
			}
			name = scanner.Text()
			name = name[1:]
			start = false
		} else {
			sequence += scanner.Text()
		}
	}
	out <- Protein{name, sequence, calculateMass(sequence)}
	close(out)
}

// map of amino acid average masses
var averageMass = map[string]float64{
	"G": 57.05177, "A": 71.07855, "S": 87.07796, "P": 97.11623,
	"V": 99.13211, "T": 101.10474, "C": 103.14464, "I": 113.15890,
	"L": 113.15890, "N": 114.10354, "D": 115.08826, "Q": 128.13032,
	"K": 128.17358, "E": 129.11504, "M": 131.19820, "H": 137.14062,
	"F": 147.17571, "R": 156.18707, "Y": 163.17512, "W": 186.21220,
}

func calculateMass(aa string) (mass float64) {
	mass = 18.000
	for _, residue := range aa {
		mass = mass + averageMass[string(residue)]
	}
	return mass / 1000
}
