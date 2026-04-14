package komagataella

import (
	//	"fmt"
	"strings"
	
	"github.com/tnaums/gobio/internal/dna"
	"github.com/tnaums/gobio/internal/protein"	
)

var alpha string = "LFETMRFPSIFTAVLFAASSALAAPVNTTTEDETAQIPAEAVIGYSDLEGDFDVAVLPFSNSTNNGLLFINTTIASIAAKEEGVSLEKREAEA"

func GetMatureProtein(d dna.DNA)  protein.Protein {
	for _, orf := range d.Orfs {
		if strings.HasPrefix(orf.AminoAcid, "LFETMRF") {
			sequence := strings.TrimPrefix(orf.AminoAcid[:len(orf.AminoAcid) -1], alpha)
			return protein.NewProtein("secreted_protein", sequence)
		}
	}
	return protein.Protein{}
}

func Hello() []byte {
	return []byte("Hello\n")
}
