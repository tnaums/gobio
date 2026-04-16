package komagataella

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tnaums/gobio/internal/dna"
	"github.com/tnaums/gobio/internal/protein"
)
// gap prefix = "FICPYFNQLNNYFET" aox1 prefix = "LFET"
var alpha string = "LFETMRFPSIFTAVLFAASSALAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"
var ost1 string = "LFETMRQVWFSWIVGLFLCFFNVSSAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"

var alphaGap string = "FICPYFNQLNNYFETMRFPSIFTAVLFAASSALAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"
var ost1Gap string = "FICPYFNQLNNYFETMRQVWFSWIVGLFLCFFNVSSAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"

//var gapCyto string = "QPRDGKVPAVAGNNSGRTHVMRLLETTRIEYKRRTPFPILVSPDPKTLNLIYLSLFQSIEQLFRNEEF"

type Komagataella struct {
	DNA     dna.DNA
	Protein protein.Protein
	SSS     string
	Promoter string
}

func (k Komagataella) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s\n", k.Protein))
	builder.WriteString(fmt.Sprintf("Secretion: %s\n", k.SSS))
	builder.WriteString(fmt.Sprintf("Plasmid size: %d\n", len(k.DNA.Parent)))
	builder.WriteString(fmt.Sprintf("Promoter: %s\n", k.Promoter))
	return builder.String()
}

func NewKomagataella(r io.Reader) (Komagataella, error) {
	dnas := dna.DNAChannelFasta(r)
	dna, ok := <-dnas
	if !ok {
		return Komagataella{}, fmt.Errorf("error retrieving DNA from reader")
	}

	protein, promoter, sss, err := getMatureProtein(dna)
	if err != nil {
		return Komagataella{}, fmt.Errorf("error getting mature protein sequence: %v", err)
		os.Exit(1)
	}
	
	return Komagataella{
		DNA:     dna,
		Protein: protein,
		SSS: sss,
		Promoter: promoter,
	}, nil
}

func getMatureProtein(d dna.DNA) (protein.Protein, string, string, error) {
	for _, orf := range d.Orfs {
		promoter := "unknown"
		if strings.HasPrefix(orf.AminoAcid, "LFETM") {
			promoter = "inducible aox1"
			if strings.HasPrefix(orf.AminoAcid, alpha) {
				sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], alpha)
				return protein.NewProtein(d.Header, sequence), promoter, "alpha", nil
			}
			if strings.HasPrefix(orf.AminoAcid, ost1) {
				sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], ost1)
				return protein.NewProtein(d.Header, sequence), promoter, "ost1", nil
			}
			sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], "LFET")
			return protein.NewProtein(d.Header, sequence), promoter, "cytoplasmic", nil
		}

		if strings.HasPrefix(orf.AminoAcid, "FICPYFNQLNNYFET") {
			promoter = "constitutive gap"
			if strings.HasPrefix(orf.AminoAcid, alphaGap) {
				sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], alphaGap)
				return protein.NewProtein(d.Header, sequence), promoter, "alpha", nil
			}
			if strings.HasPrefix(orf.AminoAcid, ost1Gap) {
				sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], ost1Gap)
				return protein.NewProtein(d.Header, sequence), promoter, "ost1", nil
			}
			sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], "FICPYFNQLNNYFET")
			return protein.NewProtein(d.Header, sequence), promoter, "cytoplasmic", nil
		}			
			
	}
	return protein.Protein{}, "", "unknown", fmt.Errorf("failed to find recombinant protein in GetMatureProtein for %s", d.Header)
}
