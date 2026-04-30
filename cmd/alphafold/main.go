package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/tnaums/gobio/internal/alphafold"
	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	// Set the uniprot id for the structure
	id := "I1S3A5"
	//id := "C7YS44"

	// Initialize client for api request
	alphafoldClient := alphafold.NewClient(15 * time.Second)

	// Returns *alphafold.AlhpafoldSummary
	resp, err := alphafoldClient.GetCIF(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// save structure to disk
	out, err := os.Create("cmd/alphafold/af.cif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer out.Close()

	io.Copy(out, resp.Body)

	file, err := os.Open("cmd/alphafold/af.cif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create protein.Protein from info in cif file
	buf := pymol.SequenceFromCIF(file)
	proteins := protein.ProteinChannelFasta(buf)
	chainA := <-proteins
	fmt.Println(chainA)
	fmt.Println()

	// use regular expression to locate catalytic serine
	r, _ := regexp.Compile("DST")
	list := r.FindStringIndex(chainA.AminoAcid)
	fmt.Println(list)

	// reset file position and generate Chainmap map[seqid]Residue
	file.Seek(0, 0)
	chainAMap := pymol.NewChainMap(file, "A")

	var motifStart int
	var motifEnd int
	
	if len(list) == 2 {
		motifStart = chainAMap[list[0]+1].IDStart
		motifEnd = chainAMap[list[1]].IDEnd

		for i := list[0] + 1; i <= list[1]; i++ {
			fmt.Printf("%v\n", chainAMap[i])
		}
	}

	// launch pymol and create StdinPipe writer to communicate with pymol
	cmd := exec.Command("pymol", "-p", "-K", "cmd/alphafold/af.cif")
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
		pymol.SelectByID(stdin, "DST", "yellow", motifStart, motifEnd, true)

	}()
	// When pymol exits, the output is captured and printed to the command line.
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)

}
