package dna

import (
	"bufio"
	"fmt"
	"os"
)

// Opens a fasta file and returns the name and a string.
// Limited to single fasta sequence files.
func FastaParser(filename string) (name, sequence string) {
	name = ""
	sequence = ""
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s", filename)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if name == "" {
			name = scanner.Text()
			name = name[1:]
		} else {
			sequence += scanner.Text()
		}
	}
	return name, sequence

}
// Returns a reversed string. A helper function to create
// complement DNA strand from coding strand.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
