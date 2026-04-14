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
	//cif := "cif/1465415.cif"
	cif := "cif/9172_0.cif"

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
	_ = <-proteins // discard chain A sequence
	chainB := <-proteins
	fmt.Println(chainB)

	// use regular expression to locate motif
	r, _ := regexp.Compile("DRSGMGQG")
	list := r.FindStringIndex(chainB.AminoAcid)
	fmt.Println(list)

	// reset file position and generate Chainmap map[seqid]Residue
	file.Seek(0, 0)
	chainBMap := pymol.NewChainMap(file, "B")
	motifStart := chainBMap[list[0] + 1].IDStart
	motifEnd := chainBMap[list[1]].IDEnd
	for i := list[0] + 1; i <= list[1]; i++ {
		fmt.Printf("%v\n", chainBMap[i])
	}

	// reset file position and generate Structure map
	// index is atom id, value is an atom struct that
	// contains all info from ATOM line in cif file
	file.Seek(0, 0)
	structure := pymol.NewStructure(file)
	for i := motifStart; i <= motifEnd; i++ {
		fmt.Printf("%v\n", structure[i])
	}

	go func() {
		defer stdin.Close()
		// change some pymol settings from default
		pymol.CustomizeCartoon(stdin)
		pymol.SetLighting(stdin)

		// Select and modify Q49,C55,H187,N208 from ChainA. These
		// residue ids were determined manually and hard-coded.
		pymol.SelectByID(stdin, "Q", "blue", 379, 387, true)
		pymol.SelectByID(stdin, "C", "red", 419, 424, true)
		pymol.SelectByID(stdin, "H", "blue", 1419, 1428, true)
		pymol.SelectByID(stdin, "N", "blue", 1587, 1594, true)

		// Select Chain B, change color
		pymol.SelectByChain(stdin, "B", "red", "B", false)

		// Select motif that was identified by regular expression pattern match.
		pymol.SelectByID(stdin, "DRSGMGQG", "blue", motifStart, motifEnd, true)
	}()

	// When pymol exits, the output is captured and printed to the command line.
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}
