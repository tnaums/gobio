// Creates a Komagatealla type from a DNA fasta file containing
// a pPICZ plasmid sequence. Prints the extracted data.
package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/komagataella"
)

func main() {
	file, err := os.Open("sequences/pTAN254.fa")
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

	fmt.Println(k)

}
