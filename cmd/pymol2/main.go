package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	cif := "cif/8VCE.cif"

	// // launch pymol and create StdinPipe writer to communicate with pymol
	// cmd := exec.Command("pymol", "-p", "-K", cif)
	// stdin, err := cmd.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	fmt.Println(chainB)
}
