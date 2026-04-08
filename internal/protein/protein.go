// Package protein provides a protein type to store protein
// sequence information
package protein

import (
	"bufio"
	"fmt"
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

// String method; Protein implements Stringer interface
// for example: fmt.Println(protein) prints 'protein' in fasta format
func (p Protein) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(">%s|%.2fkDa\n", p.Header, p.Mass))
	for idx, base := range p.AminoAcid {
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


// Create a Protein struct from header and sequence strings
func NewProtein(header, sequence string) Protein {
	return Protein{
		Header:    header,
		AminoAcid: sequence,
		Mass:      calculateMass(sequence),
	}
}

// NewProteinFromFasta creates a slice of type Protein from a fasta file
// containing one or more protein sequences. You probably want
// ProteinChannelFasta.
func NewProteinFromFasta(filename string) ([]Protein, error) {
	returnSlice := make([]Protein, 0)
	file, err := os.Open(filename)
	if err != nil {
		return []Protein{}, err
	}

	data := fastaParser(file)
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

// fastaParser reads a fasta file, extracts the sequence name from
// the header and creates a sequence string from the sequence.
// Returns a slice of strings with alternating header and sequence.
func fastaParser(r io.Reader) (data []string) {
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

// ProteinChannelFasta reads fasta sequences from an io.ReadCloser interface,
// such as an *os.File returned from os.Open(fileName). Returns channel of
// type Protein and initiates go routine that creates Proteins and adds
// to channel.
func ProteinChannelFasta(f io.Reader) <-chan Protein {
	out := make(chan Protein)
	go func() {
		//		defer f.Close()
		defer close(out)
		start := true
		var name string
		var sequence strings.Builder
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), ">") {
				if !start {
					out <- Protein{name, sequence.String(), calculateMass(sequence.String())}
					sequence.Reset()
				}
				name = scanner.Text()
				name = name[1:]
				start = false
			} else {
				sequence.WriteString(scanner.Text())
			}
		}
		out <- Protein{name, sequence.String(), calculateMass(sequence.String())}
	}()
	return out
}

// map of amino acid average masses
var averageMass = map[string]float64{
	"G": 57.05177, "A": 71.07855, "S": 87.07796, "P": 97.11623,
	"V": 99.13211, "T": 101.10474, "C": 103.14464, "I": 113.15890,
	"L": 113.15890, "N": 114.10354, "D": 115.08826, "Q": 128.13032,
	"K": 128.17358, "E": 129.11504, "M": 131.19820, "H": 137.14062,
	"F": 147.17571, "R": 156.18707, "Y": 163.17512, "W": 186.21220,
}


// returns average mass for a peptide or protein in kDa
func calculateMass(aa string) (mass float64) {
	mass = 18.000
	for _, residue := range aa {
		mass = mass + averageMass[string(residue)]
	}
	return mass / 1000
}
