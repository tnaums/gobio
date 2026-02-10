# gobio
Module gobio provides tools for reading, parsing, and analyzing DNA and protein
sequences from standard flat file formats.

## dna.go
package dna // import "github.com/tnaums/gobio/internal/dna"

Package dna provides a Dna type to store DNA sequence information and provides
simple Dna methods.

VARIABLES
```
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
}```
    GeneticCode is a map containing the standard genetic code.


FUNCTIONS
```
func FastaParser(filename string) (name, sequence string)
    The FastaParser function opens a fasta file and extracts the sequence name
    from the header and creates a sequence string from the sequence.

func ReverseComplement(parent string) (complement string)
    Function ReverseComplement takes a DNA sequence as a string and returns the
    complement DNA strand as a string.```


TYPES

type Dna struct {
	File       string
	Name       string
	Parent     string
	Complement string
}
    The Dna struct contains the File name of the flat file, the Name of the
    sequence, the Parent DNA sequence, and the Complement DNA sequence.

func NewDnaFromFasta(filename string) Dna
    NewDnaFromFasta is a function that creates a type Dna struct from a fasta
    file that contains a single fasta entry.

func NewDnaFromSequence(sequence string) Dna
    NewDNAFromSequence is a function that creates a type Dna struct from a
    sequence string.

func (d Dna) String() string
    String is a Dna method for printing the sequence in fasta format.

func (d Dna) Translate() (orfs []string)
    Translate converts DNA sequences to a slice of strings containing all
    possible open reading frames.