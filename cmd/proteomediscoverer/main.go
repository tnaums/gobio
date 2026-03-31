package main

import (
	//"bufio"
	"fmt"
	"os"
	//	"strconv"
	"time"

	//"time"

	//"strings"

	"github.com/tnaums/gobio/internal/eutils"
	"github.com/tnaums/gobio/internal/localblast"
	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/proteomediscoverer"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <file.csv>")
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

	// Get slice of records
	records, _ := proteomediscoverer.ParseCSV(file)

	// Assemble accessions for sequence retireval
	accessions := ""
	for _, record := range records {
		accessions += fmt.Sprintf("%s,", record.Accession)
	}

	// Initialize client for ncbi eutils api request
	eutilsClient := eutils.NewClient(5 * time.Second)
	// generate *http.Response from ncbi query
	resp, err := eutilsClient.EPost(accessions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Open a channel of proteins from *http.Response (io.ReadCloser)
	proteins := protein.ProteinChannelFasta(resp.Body)
	var sliceProteins []protein.Protein
	for protein := range proteins {
		sliceProteins = append(sliceProteins, protein)
	}
	fmt.Println(sliceProteins[4])

	// run localblast for selected protein against 3 databases
	proteomes := []string{"verticillioides.aa.fasta","graminearum.aa.fasta", "subglutinans.aa.fasta", "proliferatum.aa.fasta", "Vdahliae.aa.fasta", "Cgram.fasta", "Ccarb.aa.fasta"}
	for idx, proteome := range proteomes {
		fmt.Printf(" %.2d. Performing blastp against %s\n", idx, proteome)
		blast := localblast.LocalBlast(sliceProteins[4], proteome)
		localblast.PrintBlastp(blast)
	}
}
