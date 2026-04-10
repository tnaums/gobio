package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/structure2"
)

func main() {
	file, err := os.Open("cif/9172_0.cif")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	sequence := structure2.SequenceFromCIF(file)
	fmt.Println(sequence)

}
