package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/structure"
)

func main() {
	file, err := os.Open("cif/9172_0.cif")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	structure := structure.NewStructure(file)
	for idx, atom := range structure {
		fmt.Printf("%d: %#v\n\n", idx, atom)
	}


}
