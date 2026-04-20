// Module gobio provides tools biological research, with
// emphasis on protein chemistry.
//
// This program demonstrates use of localblast, protein, and
// signalp packages to analyze a fungal proteome, selecting
// secreted proteins between 16 and 19 kDa that have homologs
// in four other fungal proteomes.
package main

import (
	"fmt"
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

	// Fungal proteome file is specified from the command line
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

	// Open signalp file from jgi mycocosm. This file summarizes
	// data for all proteins that were predicted to be secreted by
	// signalp.
	//
	// https://services.healthtech.dtu.dk/services/SignalP-6.0/
	// https://mycocosm.jgi.doe.gov/mycocosm/home

	file2, err := os.Open("genomes/Fusve2/signalp.tab")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file2.Close()

	// Parse signalp data and create map
	signalPMap, _ := signalp.NewSignalPMap(file2)

	// Create go channel that returns protein.Protein from proteome sequence file
	proteins := protein.ProteinChannelFasta(file)

	for p := range proteins { // iterate over proteins that are returned from go channel
		// get protein int from p.Header and search signalp map to see if it is present
		fields := strings.Split(p.Header, "|")
		number, _ := strconv.Atoi(fields[2])
		s, ok := signalPMap[number]

		// if protein was identified by signalp as a secreted protein
		if ok {
			// create a new protein.Protein for the secreted protein, which
			// has the secretion signal sequence removed from the amino terminus
			mHeader := p.Header + "|secreted"
			mStart := s.NnCutPos
			mSequence := p.AminoAcid[mStart:]
			mature := protein.NewProtein(mHeader, mSequence) // truncated mature protein
			// select only proteins with desired mature mass
			if mature.Mass > 16 && mature.Mass < 19 {
				matches := 0
				fmt.Println()
				fmt.Println(mature)
				// Choose genomes for local blast. For
				// each secreted protein in Fusarium
				// verticillioides with mature mass in
				// selected range we will perform localblast
				// against four other fungal proteomes
				// and only keep those that have
				// homologs in all four.
				proteomes := []string{"graminearum.aa.fasta", "subglutinans.aa.fasta", "proliferatum.aa.fasta", "Ccarb.aa.fasta"}
				for _, proteome := range proteomes {
					blast := localblast.LocalBlast(p, proteome)
					qlen, _ := strconv.Atoi(blast.BlastOutputQueryLen)
					alen, _ := strconv.Atoi(blast.BlastOutputIterations.Iteration.IterationHits.Hit[0].HitHsps.Hsp[0].HspAlignLen)
					hlen, _ := strconv.Atoi(blast.BlastOutputIterations.Iteration.IterationHits.Hit[0].HitLen)

					aPercent := float64(alen) / float64(qlen)
					hSize := float64(hlen) / float64(qlen)
					// chosen criteria for defining what is a 'homolog'
					if aPercent > 0.9 && aPercent < 1.1 && hSize < 1.5 {
						matches += 1
					}
				}
				if matches == 4 {
					selected = append(selected, mature)
				}
			}
		}
	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))
	for _, protein := range selected {
		fmt.Printf("%s\n", protein)
	}

}
