package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/tnaums/gobio/internal/esmfold"
	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	// open protein fasta file
	file, err := os.Open("cmd/esmandpymol/mature.pep")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create a protein channel
	proteins := protein.ProteinChannelFasta(file)
	// get protein from channel
	p := <-proteins

	// Initialize client for api request
	esmClient := esmfold.NewClient(15 * time.Second)

	// Returns *http.Response with predicted structure in pdb format
	resp, err := esmClient.GetStructure(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// save structure to disk
	out, err := os.Create("cmd/esmandpymol/esmfold.pdb")
	if err != nil {
		panic(err)
	}
	defer out.Close() // no error handling

	io.Copy(out, resp.Body)

	file, err = os.Open("cmd/esmandpymol/esmfold.pdb")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// create protein.Protein from info in pdb file
	buf := pymol.SequenceFromPDB(file)
	proteins = protein.ProteinChannelFasta(buf)
	chainA := <-proteins
	fmt.Println(chainA)
	fmt.Println()

	// use regular expression to locate motif
	r, _ := regexp.Compile("DRSGMGQG")
	list := r.FindStringIndex(chainA.AminoAcid)
	fmt.Println(list)

	// reset file position and generate Chainmap map[seqid]Residue
	file.Seek(0, 0)
	chainAMap := pymol.NewChainMapPDB(file, "A")
	motifStart := chainAMap[list[0] + 1].IDStart
	motifEnd := chainAMap[list[1]].IDEnd
	for i := list[0] + 1; i <= list[1]; i++ {
		fmt.Printf("%v\n", chainAMap[i])
	}
	

	// launch pymol and create StdinPipe writer to communicate with pymol
	cmd := exec.Command("pymol", "-p", "-K", "cmd/esmandpymol/esmfold.pdb")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		// change some pymol settings from default
		pymol.CustomizeCartoon(stdin)
		pymol.SetLighting(stdin)

		// Set color
		pymol.SelectByChain(stdin, "chainA", "forest", "A", false)

		// Select motif that was identified by regular expression pattern match.
		pymol.SelectByID(stdin, "DRSGMGQG", "yellow", motifStart, motifEnd, true)		

	}()
	// When pymol exits, the output is captured and printed to the command line.
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)

}
