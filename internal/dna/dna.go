// Package dna provides a DNA type to store DNA
// sequence information and provides simple DNA methods.
package dna

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// The DNA struct contains the sequence header, the Parent DNA sequence,
// and the Complement DNA sequence. The Orfs slice contains all
// possible open reading frames based solely on translation.
type DNA struct {
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
func (d DNA) Translate() (orfs []Orf) {
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

// DNA.String prints the sequence of the Parent strand in fasta format.
func (d DNA) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(">%s\n", d.Header))
	for idx, base := range d.Parent {
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

// Orf.String prints the sequence of the orf in fasta format
func (o Orf) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf(">%s|Frame_%d|Length%d\n", o.Strand, o.Frame, len(o.AminoAcid)))
	for idx, base := range o.AminoAcid {
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

// DNAChannelFasta reads fasta sequences from an io.ReadCloser interface,
// such as an *os.File returned from os.Open(fileName). Returns channel of
// type DNA and initiates a go routine that creates DNAs and adds them
// to the channel.
func DNAChannelFasta(f io.ReadCloser) <-chan DNA {
	out := make(chan DNA)
	go func() {
		defer f.Close()
		defer close(out)
		start := true
		name := ""
		sequence := ""
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), ">") {
				if !start {
					newDNA := DNA{
						Header:     name,
						Parent:     sequence,
						Complement: reverseComplement(sequence),
					}
					newDNA.Orfs = newDNA.Translate()
					out <- newDNA
					sequence = ""
				}
				name = scanner.Text()
				name = name[1:]
				start = false
			} else {
				sequence += scanner.Text()
			}
		}
		newDNA := DNA{
			Header:     name,
			Parent:     sequence,
			Complement: reverseComplement(sequence),
		}
		newDNA.Orfs = newDNA.Translate()
		out <- newDNA

	}()
	return out
}

// NewDNAFromSequence is a function that creates a type DNA struct
// from a sequence string.
func NewDNAFromSequence(header, sequence string) DNA {
	newDNA := DNA{
		Header:     header,
		Parent:     sequence,
		Complement: reverseComplement(sequence),
	}
	newDNA.Orfs = newDNA.Translate()
	return newDNA
}

// NewDNAFromFasta creates a slice of type DNA from a fasta file containing
// one or more DNA sequences. You probably want
// DNACHannelFasta.
func NewDNAFromFasta(filename string) ([]DNA, error) {
	returnSlice := make([]DNA, 0)
	file, err := os.Open(filename)
	if err != nil {
		return []DNA{}, err
	}

	data := fastaParser(file)
	for i := 0; i < len(data); i = i + 2 {
		newDNA := DNA{
			Header:     data[i],
			Parent:     data[i+1],
			Complement: reverseComplement(data[i+1]),
		}
		newDNA.Orfs = newDNA.Translate()
		returnSlice = append(returnSlice, newDNA)
	}
	return returnSlice, nil
}

// FastaParser reads a fasta file, extracts the sequence name from the header,
// and creates a sequence string from the sequence. Returns a slice of strings
// with alternating header and sequence.
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

// Reverses a string. Called by reverseComplement.
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

// reverseComplement
func reverseComplement(parent string) (complement string) {

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
