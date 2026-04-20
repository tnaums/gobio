package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	cif := "cif/chita_bzcmp.cif"

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
	proteins := protein.ProteinChannelFasta(buf)
	chainA := <-proteins
	chainB := <-proteins
	fmt.Println(chainA)
	fmt.Println()
	fmt.Println(chainB)

	// use regular expression to locate polyglycine in chain A
	r, _ := regexp.Compile("G{11}")
	list := r.FindStringIndex(chainA.AminoAcid)
	fmt.Println(list)

	// reset file position and generate Chainmap map[seqid]Residue
	file.Seek(0, 0)
	chainAMap := pymol.NewChainMap(file, "A")
	polygStart := chainAMap[list[0] + 1].IDStart
	polygEnd := chainAMap[list[1]].IDEnd
	for i := list[0] + 1; i <= list[1]; i++ {
		fmt.Printf("%v\n", chainAMap[i])
	}

	// use regular expression to locate catalytic motif in chain B
	r, _ = regexp.Compile("SVSK")
	list = r.FindStringIndex(chainB.AminoAcid)
	fmt.Println(list)

	// reset file position and generate Chainmap map[seqid]Residue
	file.Seek(0, 0)
	chainBMap := pymol.NewChainMap(file, "B")
	svskStart := chainBMap[list[0] + 1].IDStart
	svskEnd := chainBMap[list[1]].IDEnd
	for i := list[0] + 1; i <= list[1]; i++ {
		fmt.Printf("%v\n", chainBMap[i])
	}
	

	go func() {
		defer stdin.Close()
		// change some pymol settings from default
		pymol.CustomizeCartoon(stdin)
		pymol.SetLighting(stdin)

		// Set colors for each chain
		pymol.SelectByChain(stdin, "chainA", "forest", "A", false)
		pymol.SelectByChain(stdin, "chainB", "red", "B", false)


		// Select motif that was identified by regular expression pattern match.
		pymol.SelectByID(stdin, "polygG", "yellow", polygStart, polygEnd, true)
		pymol.SelectByID(stdin, "SVSK", "yellow", svskStart, svskEnd, true)		
	}()

	// When pymol exits, the output is captured and printed to the command line.
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}
