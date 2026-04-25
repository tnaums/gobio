package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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

	// get a protein channel
	proteins := protein.ProteinChannelFasta(file)
	// get protein from channel
	protein := <-proteins

	// Initialize client for api request
	esmClient := esmfold.NewClient(15 * time.Second)

	// Returns predicted structure in pdb format
	resp, err := esmClient.GetStructure(protein.AminoAcid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// save file to disk
	out, err := os.Create("cmd/esmandpymol/esmfold.pdb")
	if err != nil {
		panic(err)
	}
	defer out.Close() // no error handling

	io.Copy(out, resp.Body)

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

	}()
	// When pymol exits, the output is captured and printed to the command line.
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)

}
