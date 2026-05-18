# gobio

## description

The gobio module contains a series of go packages (libraries) to
facilitate biological research with an emphasis on protein science.

## Motivation

As a protein scientist and coding enthusiast I enjoy writing computer
programs that support and empower laboratory science. One day I
discovered the go programming language. The more I learned go, the
more I wanted to code in go. I developed gobio as a way to
learn and enjoy go while building tools to support my daily
research.

## Quick Start

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.


Inside a go module:
```sh
go get github.com/tnaums/gobio
```

Import the package into your project.

```go
import "github.com/tnaums/gobio"
```

Packages like `dna` and `protein` are useful in many situations, while
others like `signalp` and `komagataella` are more specialized--useful
to scientists studying secreted fungal proteins or expressing
recombinant proteins in yeast. gobio also contains packages for
retrieving dna and protein sequences from ncbi and uniprot
databases, performing local blast searches and viewing the results,
and interacting with the pymol molecular structure viewer.

As an example, we can open a protein fasta file from disk and create a protein.Protein type. First, we open a file to create an *os.File
```go
package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/protein"
)

	// Open file to create *os.File which is an io.Reader
	file, err := os.Open("sequences/C7YS44.1.fasta")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

```
We next create a go channel. Since the function uses the io.Reader interface, an *http.Response.Body can also be used as as argument. Any io.Reader containing protein fasta sequences will work.
```go
	// Create a channel of Proteins
	proteins := protein.ChannelFromFasta(file)

```
Proteins can then be read through the channel.
```go
	protein, ok := <-proteins
	if ok {
		fmt.Println(protein)
	} else {
		fmt.Println("Protein channel is empty.")
	}
```
This prints the protein in fasta format. Since the file only had one protein sequence, a second call to the channel will not return a Protein.
```go
	protein, ok = <-proteins
	if ok {
		fmt.Println(protein)
	} else {
		fmt.Println("Protein channel is empty.")
	}	
```
```console
>sp|C7YS44.1|PGH_FUSV7|70.41kDa
MHSLSLRRLLTSVLSLCSCSSALPNQRRSNVTSHVETYYSVDGATHAEKSKALKADGYRI
VSLSSYGSPDSANYAAIWVQEEGPSFEIIHDADEATYNSWLQTWKSRGYVSTQVSATGPA
ENAVFAGVMENINVANWFQSCELENPWAFSNTTGNVDVVVKGFRMFGTPEERRYCILGHE
NVGNEQTTIQYSTPSFTVNFASTFEAETTKRFWRPSRLFLSEDHIITPSFADTSVGKWSH
AVDLTKAELKEKIETERAKGLYPIDIQGGGSGSSERFTVVFAERTSPKPRQWNVRGEITG
FEDNKAAEEEVDSIMRRFMEKNGVRQAQFAVALEGKTIAERSYTWAEDDRAIVEPDDIFL
LASVSKMFLHASIDWLVSHDMLNFSTPVYDLLGYKPADSRANDINVQHLLDHSAGYDRSM
SGDPSFMFREIAQSLPTKGAKAATLRDVIEYVVAKPLDFTPGDYSAYSNYCPMLLSYVVT
NITGVPYLDFLEKNILDGLNVRLYETAASKHTEDRIVQESKNTGQDPVHPQSAKLVPGPH
GGDGAVKEECAGTFAMAASASSLAKFIGSHAVWGTGGRVSSNRDGSLSGARAYVESRGTI
DWALTLNTREYISETEFDELRWYSLPDFLSAFPIAG

Protein channel is empty.
```

Alternately, load entire proteome.

```go
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

	// Open file to create *os.File which is an io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a channel of Proteins
	proteins := protein.ChannelFromFasta(file)

	var count int

        // iterate over the entire proteome looking form large proteins
	for protein := range proteins {
		if protein.Mass > 300 {
			fmt.Println(protein.Header)
			fmt.Println()
			count++
		}
	}
	fmt.Printf("Found %d proteins larger than 300 kDa\n", count)
```

```console
Welcome to gobio!
jgi|Fusve2|126|FVEG_00094T0
jgi|Fusve2|844|FVEG_00563T0
jgi|Fusve2|1851|FVEG_01180T0
jgi|Fusve2|4404|FVEG_09123T0
jgi|Fusve2|4405|FVEG_09123T1
jgi|Fusve2|5568|FVEG_16695T0
jgi|Fusve2|6384|FVEG_10547T0
jgi|Fusve2|6655|FVEG_10756T0
jgi|Fusve2|7062|FVEG_11086T0
jgi|Fusve2|7068|FVEG_11092T0
jgi|Fusve2|8053|FVEG_11762T0
jgi|Fusve2|8269|FVEG_11932T0
jgi|Fusve2|8523|FVEG_17172T0
jgi|Fusve2|8995|FVEG_12503T0
jgi|Fusve2|10843|FVEG_12610T0
jgi|Fusve2|13175|FVEG_15132T0
jgi|Fusve2|13845|FVEG_03249T0
jgi|Fusve2|14868|FVEG_03990T0
jgi|Fusve2|15073|FVEG_04129T0
jgi|Fusve2|15135|FVEG_15418T0
jgi|Fusve2|15947|FVEG_04724T2
jgi|Fusve2|15948|FVEG_04724T1
jgi|Fusve2|15949|FVEG_04724T0
jgi|Fusve2|16992|FVEG_05323T0
jgi|Fusve2|19341|FVEG_06977T0

Found 25 proteins larger than 300 kDa
----------------------------------------
```

## Usage
The `gobio/cmd/` directory contains example programs demonstrating how
packages work. The `main.go` files are commented and can be run
from the root directory: `go run ./cmd/demofastaprotein` or `go run
./cmd/demoeutils`

Not all ./cmd/ examples will run because the data files are not all
included in the repository--like folders of genome sequences and local
blast databases. The comments inbeded in the main.go files should,
however, explain the API.

## Contributing

### Clone the repo

```bash
git clone https://github.com/tnaums/gobio
cd gobio
```

### Run an example program

```bash
go run ./cmd/demofastadna
```

```bash
go run ./cmd/demopymol
```
requires pymol molecular structure viewer

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

# Overview of packages

Each package is contained in a single folder inside of
`gobio/internal`.

## protein

The protein package defines the Protein type. Proteins can be created by
calling `ChannelFromFasta` with an io.Reader containing one or more
sequences in fasta format.
```go
func ChannelFromFasta(f io.Reader) <-chan Protein
```

Single Protein types can be created by passing a header and sequence as
strings to `NewProtein`.
```go
func NewProtein(header, sequence string) Protein
```

## dna

The dna package defines the DNA type. DNAs can be created by calling
`ChannelFromFasta` with an io.Reader containing one or more sequences
in fasta format. The Complement sequence is automatically created as
is a slice of Orfs, which are calculated based only on translation of
all six reading frames. Some Orfs contain only a single amino acid;
most useful in a loop with size selection.
```go
func ChannelFromFasta(f io.Reader) <-chan DNA
```

Single DNA types can be created by passing a header and sequence as strings to `NewDNAFromSequence`.
```go
func NewDNAFromSequence(header, sequence string) DNA 
```

## pymol

The pymol package supports control of the pymol structure viewer from
go. It also includes functions to parse cif files into go data
structures. In the `./cmd/demopymol` example, three cif files
generated by alphafold3 predict the interaction between a plant
chitinase and three related fungal proteases that cleave it.  The
program locates peptide motifs within each protein and instructs pymol
to show these regions as stylized sticks. This allows easy analysis of
the predicted structures: in all three the polyglycine sequence is not
in the active site of the protease and the predicted structures are
hallucinations.

<img src="pictures/chita_bzcmp.png">
<img src="pictures/chita_escmp.png">
<img src="pictures/chita_fvancmp.png">

The pymol package is also useful when combined with the alphafold and
esmfold packages described below.

Package pymol was inspired by the python package 'pymolPy3':
    https://github.com/carbonscott/pymolPy3/tree/main



# komagataella
package komagataella // import "github.com/tnaums/gobio/internal/komagataella"

Package for analysis of pPICZ plasmids that are used for expression of
recombinant proteins in Komagataella pfaffii, also known as Pichia pastoris.

FUNCTIONS
```go
func GetCoding(d dna.DNA, promoter string) (dna.DNA, error)
    Extracts protein coding sequence from the plasmid sequence.
```
```go
func GetPromoter(d dna.DNA) (string, error)
    Determines promoter type, either inducible aox1, or constitutive gap.
```
```go
func GetRecombinant(dna dna.DNA) (protein.Protein, string)
    Determines protein sequence from DNA coding sequence. Also, returns
    secretion signal sequence or cytoplasmic.
```

TYPES
```go
type Komagataella struct {
	Plasmid  dna.DNA
	Promoter string
	Coding   dna.DNA
	Protein  protein.Protein
	SSS      string
}
```
```go
func NewKomagataella(r io.Reader) (Komagataella, error)
    Factory function that creates Komagataella struct from a fasta file.
```
```go
func (k Komagataella) String() string
```
