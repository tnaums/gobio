// Package dna provides a Dna type to store DNA
// sequence information and provides simple Dna methods.
package dna

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// The Dna struct contains the sequence header, the Parent DNA sequence,
// and the Complement DNA sequence. The Orfs slice contains all
// possible open reading frames based solely on translation.
type Dna struct {
	Header     string
	Parent     string
	Complement string
	Orfs       []Orf
}

// The Orf struct contains information for a possible
// open reading frame.
type Orf struct {
	Strand    string
	Frame     int
	AminoAcid string
}

// GeneticCode is a map of the standard genetic code.
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
			current = ""
		}
	}
	return orfs
}

// Dna.String prints the sequence of the Parent strand in fasta format.
func (d Dna) String() string {
	s := ">" + d.Header + "\n"
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

// Orf.String prints the sequence of the orf in fasta format
func (o Orf) String() string {
	s := fmt.Sprintf(">%s|Frame_%d|Length%d\n", o.Strand, o.Frame, len(o.AminoAcid))
	for idx, base := range o.AminoAcid {
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


// DnaPipeFasta reads fasta sequences from an io.Reader interface,
// such as an *os.File returned from os.Open(fileName) or an
// *http.Response. Returns stream of Dna structs through the
// provided go channel. Once the last Dna is sent, closes the
// channel.
func DnaPipeFasta(r io.Reader, out chan<- Dna) {
	start := true
	name := ""
	sequence := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ">") {
			if !start {
				newDna := Dna{
					Header:       name,
					Parent:     sequence,
					Complement: ReverseComplement(sequence),
				}
				newDna.Orfs = newDna.Translate()
				out <- newDna
				sequence = ""
			}
			name = scanner.Text()
			name = name[1:]
			start = false
		} else {
			sequence += scanner.Text()
		}
	}
	newDna := Dna{
		Header:       name,
		Parent:     sequence,
		Complement: ReverseComplement(sequence),
	}
	newDna.Orfs = newDna.Translate()
	out <- newDna
	close(out)
}

// NewDNAFromSequence is a function that creates a type Dna struct
// from a sequence string.
func NewDnaFromSequence(sequence string) Dna {
	newDna := Dna{Parent: sequence,
		Complement: ReverseComplement(sequence),
	}
	newDna.Orfs = newDna.Translate()
	return newDna
}

// NewDnaFromFasta creates a type Dna struct from a fasta file containing
// a single DNA sequence
func NewDnaFromFasta(filename string) (Dna, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Dna{}, err
	}

	header, sequence := FastaParser(file)
	newDna := Dna{
		Header:       header,
		Parent:     sequence,
		Complement: ReverseComplement(sequence),
	}
	newDna.Orfs = newDna.Translate()
	return newDna, nil
}

// FastaParser reads a fasta file, extracts the sequence name from the header,
// and creates a sequence string from the sequence.
func FastaParser(r io.Reader) (name, sequence string) {
	name = ""
	sequence = ""
	scanner := bufio.NewScanner(r)
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

// Reverses a string. Called by ReverseComplement.
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

// ReverseComplement creates a complement DNA strand.
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
