package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/tnaums/gobio/internal/dna"
)

func main() {
	fmt.Println("Welcome to gobio!")
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <sequence.fa> <dna_search_string>")
	}

	fileName := os.Args[1]
	subseq := os.Args[2]
	dnaStruct := dna.NewDnaFromFasta(fileName)

	fmt.Println()
	fmt.Println(" parent strand")
	fmt.Println("--------------------")
	x := longestMatch(dnaStruct.Parent, subseq)
	fmt.Println(x)
	fmt.Println()
	fmt.Println(" complement strand")
	fmt.Println("--------------------")
	y := longestMatch(dnaStruct.Complement, subseq)
	fmt.Println(y)

	fmt.Println(dnaStruct)

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
