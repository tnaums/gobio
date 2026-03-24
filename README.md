# dna.go
package dna // import "github.com/tnaums/gobio/internal/dna"

Package dna provides a DNA type to store DNA sequence information and provides
simple DNA methods.

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
func DNAChannelFasta(f io.ReadCloser) <-chan DNA
    DNAChannelFasta reads fasta sequences from an io.ReadCloser interface, such
    as an *os.File returned from os.Open(fileName). Returns channel of type DNA
    and initiates a go routine that creates DNAs and adds them to the channel.
```

TYPES
```go
type DNA struct {
	Header     string
	Parent     string
	Complement string
	Orfs       []Orf
}
    The DNA struct contains the sequence header, the Parent DNA sequence,
    and the Complement DNA sequence. The Orfs slice contains all possible open
    reading frames based solely on translation.
```
```go
func NewDNAFromFasta(filename string) ([]DNA, error)
    NewDNAFromFasta creates a slice of type DNA from a fasta file containing one
    or more DNA sequences.
```
```go
func NewDNAFromSequence(header, sequence string) DNA
    NewDNAFromSequence is a function that creates a type DNA struct from a
    sequence string.
```
```go
func (d DNA) String() string
    DNA.String prints the sequence of the Parent strand in fasta format.
```
```go
func (d DNA) Translate() (orfs []Orf)
    Translate converts DNA sequences to a slice of type Orf containing all
    possible open reading frames.
```

```go
type Orf struct {
	Strand    string
	Frame     int
	AminoAcid string
}

    The Orf struct contains information for a possible open reading frame.
```    
```go
func (o Orf) String() string
    Orf.String prints the sequence of the orf in fasta format
```

# protein.go
package protein // import "github.com/tnaums/gobio/internal/protein"

Package protein provides a protein type to store protein sequence information

FUNCTIONS
```go
func ProteinChannelFasta(f io.ReadCloser) <-chan Protein
    ProteinChannelFasta reads fasta sequences from an io.ReadCloser interface,
    such as an *os.File returned from os.Open(fileName). Returns channel of type
    Protein and initiates go routine that creates Proteins and adds to channel.
```

TYPES
```go
type Protein struct {
	Header    string
	AminoAcid string
	Mass      float64
}
    Contains header and amino acid sequence, parsed from fasta file. Mass can be
    calculated from AminoAcid by calling calculateMass(aaSequence)
```
```go
func NewProtein(header, sequence string) Protein
    Create a Protein struct from header and sequence strings
```
```go
func NewProteinFromFasta(filename string) ([]Protein, error)
    NewProteinFromFasta creates a slice of type Protein from a fasta file
    containing one or more protein sequences.
```
```go
func (p Protein) String() string
    String method; Protein can be used as Stringer interface;
    for example: fmt.Println(protein) prints 'protein' in fasta format
```
# DNA Tutorial
```go
// This script demonstrates how to work with dna sequences
// from fasta files containing multiple entries.
package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	fmt.Println("Welcome to gobio!")

	// Get filename from command line
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Open file to create *os.File which implements io.ReadCloser
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Create channel of DNA from io.ReadCloser interface
	dnas := dna.DNAChannelFasta(file)

	// Retrieve first sequence and print
	first :=  <- dnas
	fmt.Println(first)

	// Retrieve second sequence and print
	second := <- dnas
	fmt.Println(second)

	// Iterate over orfs and print if over 100 amino acids
	for idx, orf := range second.Orfs {
		if len(orf.AminoAcid) > 100 {
			fmt.Println(idx)
			fmt.Println(orf)
			fmt.Println()
		}
	}

	// Iterate over remaining dna sequences and print the
	// header line and sequence length
	for d := range dnas {
		fmt.Println(d.Header)
		fmt.Printf("Length: %d\n", len(d.Parent))
	}
}
```
# Protein Tutorial
```go
// Demonstrates use of protein.ProteinChannelFasta for
// reading proteins from a fasta file containing one
// or more sequences.
package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/protein"
)

func main() {
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Open file to create *os.File which is an io.ReadCloser
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a channel of Proteins
	proteins := protein.ProteinChannelFasta(file)

        // Read first protein from channel
	protein, ok := <-proteins
	if ok {
		fmt.Println(protein)
	} else {
		fmt.Println("Protein channel is empty.")
	}

        // Read second protein
	protein, ok = <-proteins
	if ok {
		fmt.Println(protein)
	} else {
		fmt.Println("Protein channel is empty.")
	}	

        // Read all remaining proteins. Print header of proteins over 200 kDa
	var count int
	for protein := range proteins {
		if protein.Mass > 200 {
			fmt.Println(protein.Header)
			fmt.Println()
			count++
		}
	}
	fmt.Printf("Found %d proteins larger than 20 kDa\n", count)
}
```
# eutils Tutorial
```go
// Demonstrates use of eutils.EPost for retrieving protein fasta
// sequences from NCBI. The response body is then sent to
// protein.ProteinChannelFasta where the seqeunces are returned as
// protein.Protein type through a go channel.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/protein"
)

func main() {
	// Initialize client for api request
	eutilsClient := eutils.NewClient(5 * time.Second)
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost("AIZ65945.1,QIR83317.1,194680922,50978626,28558982,9507199,6678417")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body) 

	// Print first protein
	fmt.Println(<-proteins)
	fmt.Println()

	// Print sequence from second protein
	p2 := <-proteins
	fmt.Println(p2.AminoAcid)
	fmt.Println()

	// For remaining proteins, print header, mass, sequence length
	for p := range proteins { 
		fmt.Printf(">%s|%.2fkDa|%daa", p.Header, p.Mass, len(p.AminoAcid))
		fmt.Println()
	}
}
```