package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	//cif := "cif/1465415.cif"
	cif := "cmd/1BS9/1BS9.cif"
	//cif := "cif/mutant2.cif"
	// cif := "cif/cocca.cif"
	//cif := "cif/1465415.cif"

	// launch pymol and create StdinPipe writer to communicate with pymol
	cmd := exec.Command("pymol", "-p", "-K", cif)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// open file to find motif
	file, err := os.Open(cif)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// create protein.Protein from info in cif file
	buf := pymol.SequenceFromCIF(file)
	proteins := protein.ChannelFromFasta(buf)
	chainA := <-proteins

	fmt.Println(chainA)

	/*
	>ChainA|20.65kDa
	SCPAIHVFGARETTASPGYGSSSTVVNGVLSAYPGSTAEAINYPACGGQSSCGGASYSSS
	VAQGIAAVASAVNSFNSQCPSTKIVLVGYSQGGEIMDVALCGGGDPNQGYTNTAVQLSSS
	AVNMVKAAIFMGDPMFRAGLSYEVGTCAAGGFDQRPAGFSCPSAAKIKSYCDASDPYCCN
	GSNAATHQGYGSEYGSQALAFVKSKLG
	*/

	go func() {
		defer stdin.Close()
		// change some pymol settings from default
		pymol.CustomizeCartoon(stdin)
		pymol.SetLighting(stdin)

		// Select Chains and set color
		pymol.SelectByChain(stdin, "chainA", "forest", "A", false)

		pymol.SelectNone(stdin)
	}()

	// When pymol exits, the output is captured and printed to the command line.
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}
