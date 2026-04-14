package komagataella

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tnaums/gobio/internal/dna"
	"github.com/tnaums/gobio/internal/protein"
)

var alpha string = "LFETMRFPSIFTAVLFAASSALAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"
var ost1 string = "LFETMRQVWFSWIVGLFLCFFNVSSAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"

type SecretionSignalSequence int

const (
	alphasignal SecretionSignalSequence = iota
	ost1signal
	cytoplasmic
)

type Komagataella struct {
	DNA     dna.DNA
	Protein protein.Protein
	SSS     SecretionSignalSequence
}

func NewKomagataella(r io.Reader) (Komagataella, error) {
	dnas := dna.DNAChannelFasta(r)
	dna, ok := <-dnas
	if !ok {
		return Komagataella{}, fmt.Errorf("error retrieving DNA from reader")
	}

	protein, s, err := GetMatureProtein(dna)
	if err != nil {
		return Komagataella{}, fmt.Errorf("error getting mature protein sequence: %v", err)
		os.Exit(1)
	}
	return Komagataella{
		DNA:     dna,
		Protein: protein,
		SSS: s,
	}, nil
}

func GetMatureProtein(d dna.DNA) (protein.Protein, SecretionSignalSequence, error) {
	for _, orf := range d.Orfs {
		if strings.HasPrefix(orf.AminoAcid, "LFETM") {
			if strings.HasPrefix(orf.AminoAcid, alpha) {
				sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], alpha)
				return protein.NewProtein(d.Header, sequence), alphasignal, nil
			}
			if strings.HasPrefix(orf.AminoAcid, ost1) {
				sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], ost1)
				return protein.NewProtein(d.Header, sequence), ost1signal, nil
			}
			sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid)-1], "LFET")
			return protein.NewProtein(d.Header, sequence), cytoplasmic, nil
		}
	}
	return protein.Protein{}, alphasignal, fmt.Errorf("failed to find recombinant protein in GetMatureProtein for %s", d.Header)
}
