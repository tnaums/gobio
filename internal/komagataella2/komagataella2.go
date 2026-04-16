package komagataella2

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/tnaums/gobio/internal/dna"
	"github.com/tnaums/gobio/internal/protein"
)

var alpha string = "MRFPSIFTAVLFAASSALAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"
var ost1 string = "MRQVWFSWIVGLFLCFFNVSSAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"

type Komagataella struct {
	Plasmid  dna.DNA
	Promoter string
	Coding   dna.DNA
	Protein  protein.Protein
	SSS      string
}

func NewKomagataella(r io.Reader) (Komagataella, error) {
	dnas := dna.DNAChannelFasta(r)
	dna, ok := <-dnas
	if !ok {
		return Komagataella{}, fmt.Errorf("error retrieving DNA from reader")
	}
	promoter, _ := GetPromoter(dna)
	coding, _ := GetCoding(dna, promoter)
	protein, sss := GetRecombinant(coding)

	return Komagataella{
		Plasmid:  dna,
		Promoter: promoter,
		Coding:   coding,
		Protein:  protein,
		SSS:      sss,
	}, nil
}

func GetPromoter(d dna.DNA) (string, error) {
	aox1, _ := regexp.Compile("AGATCTAACATC.{916}TTATTCGAAACG")
	gap, _ := regexp.Compile("AGATCTTTTTTG.{459}TTGAACAACTAT")
	if aox1.MatchString(d.Parent) {
		return "aox1", nil
	}
	if gap.MatchString(d.Parent) {
		return "gap", nil
	}
	return "unknown", fmt.Errorf("promoter not found")
}

func GetCoding(d dna.DNA, promoter string) (dna.DNA, error) {
	aox1, _ := regexp.Compile("TTATTCGAAACG(.*)GTTTGTAGCCTT")
	gap, _ := regexp.Compile("TATTTCGAAACG(.*)GTTTTAGCCTTA")
	// Ensure in-frame start at ATG. some GAP cytoplasmic
	// sequences have extra nucleotides
	startAtATG, _ := regexp.Compile("ATG.*")

	if promoter == "aox1" {
		codingString := aox1.FindStringSubmatch(d.Parent)
		codingDNA := dna.NewDNAFromSequence("coding", codingString[1])
		return codingDNA, nil
	}
	if promoter == "gap" {
		codingString := gap.FindStringSubmatch(d.Parent)
		// Some gap cytoplasmic have extra 5' nucleotides
		adjusted := startAtATG.FindString(codingString[1])
		codingDNA := dna.NewDNAFromSequence("coding", adjusted)
		return codingDNA, nil
	}
	return dna.DNA{}, nil
}

func GetRecombinant(dna dna.DNA) (protein.Protein, string) {
	proteinString := dna.Orfs[0].AminoAcid
	if strings.HasPrefix(proteinString, alpha) {
		sequence := strings.TrimPrefix(proteinString[:len(proteinString)-1], alpha)
		return protein.NewProtein("alpha|mature", sequence), "alpha"
	}
	if strings.HasPrefix(proteinString, ost1) {
		sequence := strings.TrimPrefix(proteinString[:len(proteinString)-1], ost1)
		return protein.NewProtein("ost1|mature", sequence), "ost1"
	}
	return protein.NewProtein("cytoplasmic", proteinString[:len(proteinString)-1]), "cytoplasmic"
}
