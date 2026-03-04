# dna.go
package dna // import "github.com/tnaums/gobio/internal/dna"

Package dna provides a Dna type to store DNA sequence information and provides
simple Dna methods.

VARIABLES
```go
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
```
    GeneticCode is a map of the standard genetic code.


FUNCTIONS
```go
func FastaParser(r io.Reader) (name, sequence string)
    FastaParser reads a fasta file, extracts the sequence name from the header,
    and creates a sequence string from the sequence.

func ReverseComplement(parent string) (complement string)
    ReverseComplement creates a complement DNA strand.
```

TYPES
```go
type Dna struct {
	File       string
	Name       string
	Parent     string
	Complement string
	Orfs       []Orf
}
    The Dna struct contains the File name of a fasta file (if present) ,
    the Name of the sequence (if present), the Parent DNA sequence, and the
    Complement DNA sequence. The Orfs slice contains all possible open reading
    frames based solely on translation.

func NewDnaFromFasta(filename string) (Dna, error)
    NewDnaFromFasta creates a type Dna struct from a fasta file containing a
    single DNA sequence

func NewDnaFromSequence(sequence string) Dna
    NewDNAFromSequence is a function that creates a type Dna struct from a
    sequence string.

func (d Dna) String() string
    Dna.String prints the sequence of the Parent strand in fasta format.

func (d Dna) Translate() (orfs []Orf)
    Translate converts DNA sequences to a slice of type Orf containing all
    possible open reading frames.
```

```go
type Orf struct {
	Strand    string
	Frame     int
	AminoAcid string
}
```
    The Orf struct contains information for a possible open reading frame.
    
```go
func (o Orf) String() string
    Orf.String prints the sequence of the orf in fasta format
```

# protein.go
package protein // import "github.com/tnaums/gobio/internal/protein"

Package protein provides a protein type to store protein sequence information

FUNCTIONS
```go
func FastaParser(r io.Reader) (data []string)
    FastaParser reads a fasta file, extracts the sequence name from the header
    and creates a sequence string from the sequence. Returns a slice of strings
    with alternating header and sequence.

func ProteinPipeFasta(r io.Reader, out chan<- Protein)
    ProteinPipeFasta reads fasta sequences from an io.Reader interface, such
    as an *os.File returned from os.Open(fileName). Returns stream of Protein
    structs through the provided go channel. Once the last Protein is sent,
    closes the channel.
```

TYPES
```go
type Protein struct {
	Header    string
	AminoAcid string
}
    Contains header and amino acid sequence, parsed from fasta file.

func NewProteinFromFasta(filename string) ([]Protein, error)
    NewProteinFromFasta creates a slice of type Protein from a fasta file
    containing one or more protein sequences.
```