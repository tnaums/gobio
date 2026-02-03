package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	fmt.Println("Welcome to gobio!")
	fileName := os.Args[1]
	subseq := os.Args[2]
	_, seqSequence := dna.FastaParser(fileName)
	revComp := reverseCompliment(seqSequence)

	fmt.Println()
	fmt.Println(" parent strand")
	fmt.Println("--------------------")
	x := longestMatch(seqSequence, subseq)
	fmt.Println(x)
	fmt.Println()
	fmt.Println(" complement strand")
	fmt.Println("--------------------")
	y := longestMatch(revComp, subseq)
	fmt.Println(y)

}

func longestMatch(seq, subseq string) int {
	sub := ""
	longest := 0
	for idx, s := range subseq {
		sub += string(s)
		b := strings.Contains(seq, sub)
		if b {
			longest++
			fmt.Printf(" %d. %s: %v\n", idx+1, sub, b)
		} else {
			break
		}
	}
	return longest
}

func reverseCompliment(seq string) string {
	rc := ""
	reverseSeq := reverse(seq)
	for _, base := range reverseSeq {
		if base == 'A' {
			rc += string('T')
		} else if base == 'C' {
			rc += string('G')
		} else if base == 'G' {
			rc += string('C')
		} else if base == 'T' {
			rc += string('A')
		}
	}
	return rc
}

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
