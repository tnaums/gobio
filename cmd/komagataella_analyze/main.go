package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/komagataella"
)

func main() {

	plasmids := []string{"sequences/pTAN288.fa", "sequences/pTAN303.fa", "sequences/pTAN309.fa"}
	for _, plasmid := range plasmids {
		file, err := os.Open(plasmid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()

		k, err := komagataella.NewKomagataella(file)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", k.Protein)
	}
}
