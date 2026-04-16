package main

import (
	"fmt"
	"os"

	"github.com/tnaums/gobio/internal/komagataella2"
)

func main() {
	file, err := os.Open("sequences/pTAN254.fa")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	k, err := komagataella2.NewKomagataella(file)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(k.Protein)
	fmt.Println(k.Promoter)

}
