// Package dna provides a Dna type to store DNA
// sequence information and provides simple Dna methods.
package dna

import (
	"bufio"
	"fmt"
	"os"
)

// The Dna struct contains the File name of a fasta file (if present)
// , the Name of the sequence (if present), the Parent DNA sequence,
// and the Complement DNA sequence. The Orfs slice contains all
// possible open reading frames based solely on translation.
type Dna struct {
	File       string
	Name       string
	Parent     string
	Complement string
	Orfs []Orf
}

// The Orf struct contains information for a possible
// open reading frame.
type Orf struct {
	Strand    string
	Frame     int
	AminoAcid string
}

// GeneticCode is a map containing the standard genetic
// code.
var GeneticCode = map[string]byte{
	"TTT": 'F', "TTC": 'F', "TTG": 'L', "TTA": 'L',
	"TCT": 'S', "TCC": 'S', "TCA": 'S', "TCG": 'S',
	"TAT": 'Y', "TAC": 'Y', "TAG": '*', "TAA": '*',
	"TGT": 'C', "TGC": 'C', "TGG": 'W', "TGA": '*',
	"CTT": 'L', "CTC": 'L', "CTG": 'L', "CTA": 'L',
	"CCT": 'P', "CCC": 'P', "CCA": 'P', "CCG": 'P',
	"CAT": 'H', "CAC": 'H', "CAG": 'Q', "CAA": 'Q',
	"CGT": 'R', "CGC": 'R', "CGG": 'R', "CGA": 'R',
	"ATT": 'I', "ATC": 'I', "ATG": 'M', "ATA": 'I',
	"ACT": 'T', "ACC": 'T', "ACA": 'T', "ACG": 'T',
	"AAT": 'N', "AAC": 'N', "AAG": 'K', "AAA": 'K',
	"AGT": 'S', "AGC": 'S', "AGG": 'R', "AGA": 'R',
	"GTT": 'V', "GTC": 'V', "GTG": 'V', "GTA": 'V',
	"GCT": 'A', "GCC": 'A', "GCA": 'A', "GCG": 'A',
	"GAT": 'D', "GAC": 'D', "GAG": 'E', "GAA": 'E',
	"GGT": 'G', "GGC": 'G', "GGG": 'G', "GGA": 'G',
}

// Translate converts DNA sequences to a slice of type Orf
// containing all possible open reading frames.
func (d Dna) Translate() (orfs []Orf) {
	orfs = []Orf{}
	current := ""
	complementFlag := true
	sequences := []string{d.Parent, d.Complement}
	strand := ""
	for _, sequence := range sequences {
		if complementFlag {
			complementFlag = false
			strand = "Parent"
		} else {
			complementFlag = true
			strand = "Complement"
		}
		for j := 0; j < 3; j++ { // j is the start index for each frame
			for i := j; i < len(sequence)-2; i = i + 3 {
				first := string(sequence[i])
				second := string(sequence[i+1])
				third := string(sequence[i+2])
				codon := first + second + third
				aa := string(GeneticCode[codon])
				current += aa
				if aa == "*" {
					newOrf := Orf{
						Strand:    strand,
						Frame:     j + 1,
						AminoAcid: current,
					}
					orfs = append(orfs, newOrf)
					current = ""
					continue
				}
			}
			newOrf := Orf{
				Strand:    strand,
				Frame:     j + 1,
				AminoAcid: current,
			}
			orfs = append(orfs, newOrf)
		}
	}
	fmt.Printf("number of orfs: %d\n\n", len(orfs))
	return orfs
}

// String is a Dna method for printing the sequence in
// fasta format.
func (d Dna) String() string {
	s := ">" + d.Name + "\n"
	for idx, base := range d.Parent {
		if idx == 0 {
			s += string(base)
			continue
		}
		if idx%60 == 0 {
			s += "\n"
			s += string(base)
			continue
		}
		s += string(base)

	}
	return s
}

// NewDNAFromSequence is a function that creates
// a type Dna struct from a sequence string.
func NewDnaFromSequence(sequence string) Dna {
	newDna :=  Dna{Parent: sequence,
		Complement: ReverseComplement(sequence),
	}
	newDna.Orfs = newDna.Translate()
	return newDna
}

// NewDnaFromFasta is a function that creates a
// type Dna struct from a fasta file that contains
// a single fasta entry.
func NewDnaFromFasta(filename string) Dna {
	header, sequence := FastaParser(filename)
	newDna := Dna{
		File:       filename,
		Name:       header,
		Parent:     sequence,
		Complement: ReverseComplement(sequence),
	}
	newDna.Orfs = newDna.Translate()
	return newDna
}

// The FastaParser function opens a fasta file, extracts
// the sequence name from the header, and creates a sequence
// string from the sequence.
func FastaParser(filename string) (name, sequence string) {
	name = ""
	sequence = ""
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s", filename)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if name == "" {
			name = scanner.Text()
			name = name[1:]
		} else {
			sequence += scanner.Text()
		}
	}
	return name, sequence

}

// Returns a reversed string. A helper function called by
// func ReverseComplement.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

// Function ReverseComplement takes a DNA sequence as a
// string and returns the complement DNA strand as a string.
func ReverseComplement(parent string) (complement string) {

	reverseSeq := reverse(parent)
	for _, base := range reverseSeq {
		if base == 'A' {
			complement += string('T')
		} else if base == 'C' {
			complement += string('G')
		} else if base == 'G' {
			complement += string('C')
		} else if base == 'T' {
			complement += string('A')
		}
	}
	return complement
}
