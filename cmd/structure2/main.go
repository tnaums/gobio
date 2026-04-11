package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/tnaums/gobio/internal/protein"
	"github.com/tnaums/gobio/internal/pymol"
)

func main() {
	file, err := os.Open("cif/9172_0.cif")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	buf := pymol.SequenceFromCIF(file)
	proteins := protein.ProteinChannelFasta(buf)
	_ = <-proteins
	chainB := <-proteins
	fmt.Println(chainB)

	r, _ := regexp.Compile("DRSG(M)GQG")
	list := r.FindStringIndex(chainB.AminoAcid)
	fmt.Println(list)
	file.Seek(0, 0)
	chainBMap := pymol.NewChainMap(file, "B")
	for i := list[0] + 1; i <= list[1]; i++ {
		fmt.Printf("%v\n", chainBMap[i])
	}

}
