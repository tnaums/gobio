// Module gobio provides tools for reading
// and analyzing DNA sequences from fasta files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tnaums/gobio/internal/protein"
)

func main() {
	// slice to hold proteins that fit criteria
	selected := make([]protein.Protein, 0)

	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <sequence.fa>")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// Open file to create *os.File which implements io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Create go channel
	proteins := protein.ProteinChannelFasta(file)

	for p := range proteins { // iterate over proteins that are returned from go channel
		selected = append(selected, p)
	}
	fmt.Printf("Number of selected proteins is: %d\n", len(selected))
	fmt.Println("----------------------------------------\n")

}

// confirm displays a prompt `s` to the user and returns a bool indicating yes / no
// If the lowercased, trimmed input begins with anything other than 'y', it returns false
// It accepts an int `tries` representing the number of attempts before returning false
func confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}
