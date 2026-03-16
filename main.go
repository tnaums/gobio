// Module gobio provides tools for reading
// and analyzing DNA sequences from fasta files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tnaums/gobio/internal/localblast"
	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/signalp"
)

func main() {
	// slice to hold proteins that fit criteria
	selected := make([]protein.Protein, 0)
	
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
		os.Exit(1)		
	}
	fileName := os.Args[1]

	// Open file to create *os.File which implements io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Open signalp file from jgi mycocosm
	file2, err := os.Open("genomes/Fusve2/signalp.tab")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file2.Close()

	// Parse signalp data and create map
	signalPMap, _ := signalp.NewSignalPMap(file2)

	// Create go 
	proteins := protein.ProteinChannelFasta(file)

	for p := range proteins { // iterate over proteins that are returned from go channel
		// get protein int from p.Header
		fields := strings.Split(p.Header, "|")
		number, _ := strconv.Atoi(fields[2])
		s, ok := signalPMap[number]

		// if protein identified by signalp
		if ok { 
			mHeader := p.Header + "|secreted"
			mStart := s.NnCutPos
			mSequence := p.AminoAcid[mStart:]
			mature := protein.NewProtein(mHeader, mSequence) // truncated mature protein
			if mature.Mass > 16 && mature.Mass < 19 {
				matches := 0
				fmt.Println()
				fmt.Println(mature)
				proteomes := []string{"graminearum.aa.fasta", "subglutinans.aa.fasta", "proliferatum.aa.fasta"}
				for _, proteome := range proteomes {
					blast := localblast.LocalBlast(p, proteome)
					qlen, _ := strconv.Atoi(blast.BlastOutputQueryLen)
					alen, _ := strconv.Atoi(blast.BlastOutputIterations.Iteration.IterationHits.Hit[0].HitHsps.Hsp[0].HspAlignLen)
					hlen, _ := strconv.Atoi(blast.BlastOutputIterations.Iteration.IterationHits.Hit[0].HitLen)
					fmt.Printf("query length is: %d\n", qlen)
					fmt.Printf("hit length is: %d\n", hlen)
					fmt.Printf("alignment length is: %d\n", alen)
					aPercent := float64(alen) / float64(qlen)
					hSize := float64(hlen) / float64(qlen)
					if  aPercent > 0.9 && aPercent < 1.1 && hSize < 1.5 {
						matches += 1
					} 
				}
				if matches == 3 {selected = append(selected, mature)}
			}
		}
	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))
	for _, protein := range selected {
		fmt.Printf("%s\n", protein)
	}
	fmt.Println("----------------------------------------\n")

}


// confirm displays a prompt `s` to the user and returns a bool indicating yes / no
// If the lowercased, trimmed input begins with anything other than 'y', it returns false
// It accepts an int `tries` representing the number of attempts before returning false
func confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}
