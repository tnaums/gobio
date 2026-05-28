// Package protein provides a protein representation and methods to create
// them from sequence files.
package protein

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Protein is a protein representation.  Mass is automatically
// calculated at creation. Peptides is a representation of
// tryptic peptide that can be created by calling
// protein.CreateTrypticPeptides().
type Protein struct {
	Header    string
	AminoAcid string
	Mass      float64
	Peptides  []Tryptic
}

type Tryptic struct {
	Sequence   string
}

// CreateTrypticPeptides fills the Peptides field of a Protein.
func (p *Protein) CreateTrypticPeptides() {
	results := make([]Tryptic, 0)
	sequence := p.AminoAcid
	for {
		if len(sequence) == 0 {
			break
		}
		indexArgenine := strings.Index(sequence, "R")
		indexLysine := strings.Index(sequence, "K")
		if indexArgenine == -1 && indexLysine == -1 {
			peptide := Tryptic{
				Sequence: sequence,
			}
			results = append(results, peptide)
			break
		}
		if indexArgenine == -1 {
			// Just cut at lysine
			peptide := Tryptic{
				Sequence: sequence[:indexLysine+1],
			}
			results = append(results, peptide)
			sequence = sequence[indexLysine+1:]
			continue
		}
		if indexLysine == -1 {
			// Just cut at argenine
			peptide := Tryptic{
				Sequence: sequence[:indexArgenine+1],
			}
			results = append(results, peptide)
			sequence = sequence[indexArgenine+1:]
			continue
		}
		if indexArgenine < indexLysine {
			peptide := Tryptic{
				Sequence: sequence[:indexArgenine+1],
			}
			results = append(results, peptide)
			sequence = sequence[indexArgenine+1:]
			continue
		}
		// Lysine must be the next site
		peptide := Tryptic{
			Sequence: sequence[:indexLysine+1],
		}
		results = append(results, peptide)
		sequence = sequence[indexLysine+1:]
	}
	p.Peptides = results
}

// String prints Protein in fasta format.
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

// NewFromSequence creates and returns a single protein representation
// from header and sequence strings.
func NewFromSequence(header, sequence string) Protein {
	return Protein{
		Header:    header,
		AminoAcid: sequence,
		Mass:      calculateMass(sequence),
	}
}

// ChannelFromFasta reads fasta sequences from an io.Reader
// interface.  Returns channel of type Protein and initiates go
// routine that creates and adds Proteins to channel.
func ChannelFromFasta(f io.Reader) <-chan Protein {
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
					out <- Protein{
						Header: name,
						AminoAcid: sequence.String(),
						Mass: calculateMass(sequence.String()),
					}
					sequence.Reset()
				}
				name = scanner.Text()
				name = name[1:]
				start = false
			} else {
				sequence.WriteString(scanner.Text())
			}
		}
		out <- Protein{
			Header: name,
			AminoAcid: sequence.String(),
			Mass: calculateMass(sequence.String()),
		}
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

// calculateMass returns average mass for a peptide or protein in kDa
func calculateMass(aa string) (mass float64) {
	mass = 18.000
	for _, residue := range aa {
		mass = mass + averageMass[string(residue)]
	}
	return mass / 1000
}
