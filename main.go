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
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
	}
	fileName := os.Args[1]
	selected := make([]protein.Protein, 0)

	// Create protein pipe from proteome fasta file
	proteins := make(chan protein.Protein)
	// Open file to create *os.File which implements io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	file2, err := os.Open("genomes/Fusve2/signalp.tab")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file2.Close()

	signalPMap, _ := signalp.NewSignalPMap(file2)

	// Create go routine with opened fasta file and go channel
	go protein.ProteinPipeFasta(file, proteins)

	for p := range proteins { // iterate over proteins that are returned from go channel
		// get protein int from p.Header
		fields := strings.Split(p.Header, "|")
		number, _ := strconv.Atoi(fields[2])
		s, ok := signalPMap[number]
		if ok { // protein is in the signalp list
			mHeader := p.Header + "|secreted"
			mStart := s.NnCutPos
			mSequence := p.AminoAcid[mStart:]
			mature := protein.NewProtein(mHeader, mSequence)
			if mature.Mass > 16 && mature.Mass < 19 {
				fmt.Println()
				fmt.Println(mature)				
				blast := localblast.LocalBlast(p)
				localblast.ParseBlastp(blast)
				resp := confirm("Keep this protein", 2)
				if resp {
					selected = append(selected, p)

				}
			}
		}
	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))
	fmt.Println("----------------------------------------\n")


	// // Initialize client for api request
	// eutilsClient := eutils.NewClient(50 * time.Second)
	// // generate *http.Response from ncbi query
	// resp, err := eutilsClient.EPost()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer resp.Body.Close()

	// // Create a new go channel
	// proteins = make(chan protein.Protein)
	// // start go routine with http response body as io.Reader and proteins channel
	// go protein.ProteinPipeFasta(resp.Body, proteins)

	// for p := range proteins { // iterate over proteins returned from go channel
	// 	fmt.Println(p)
	// 	fmt.Println()
	// }
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
