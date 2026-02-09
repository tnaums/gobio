package dna

import (
	"bufio"
	"fmt"
	"os"
)

type Dna struct {
	File       string
	Name       string
	Parent     string
	Complement string
}

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

func (d Dna) Translate() {
	orfs := []string{}
	current := ""
	fwdSequence := d.Parent
	for j := 0; j < 3; j++ {
		for i := j; i < len(fwdSequence)-2; i = i + 3 {
			first := string(fwdSequence[i])
			second := string(fwdSequence[i+1])
			third := string(fwdSequence[i+2])
			codon := first + second + third
			aa := string(GeneticCode[codon])
			current += aa
			if aa == "*" {
				orfs = append(orfs, current)
				current = ""
				continue
			}
		}
	}
	orfs = append(orfs, current)
	fmt.Printf("number of orfs: %d", len(orfs))
}

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

// Creates a type Dna struct from a sequence
func NewDnaFromSequence(sequence string) Dna {
	return Dna{Parent: sequence,
		Complement: ReverseComplement(sequence),
	}
}

// Creates a type Dna struct from a single fasta file
func NewDnaFromFasta(filename string) Dna {
	header, sequence := FastaParser(filename)
	return Dna{
		File:       filename,
		Name:       header,
		Parent:     sequence,
		Complement: ReverseComplement(sequence),
	}
}

// Opens a fasta file and returns the name and a string.
// Limited to single fasta sequence files.
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

// Returns a reversed string. A helper function to create
// complement DNA strand from coding strand.
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

// Takes a parent DNA strand and returns the complement strand.
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
