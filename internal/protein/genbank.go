package protein

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"strings"
)

type gbState int

const (
	gbStateSTART gbState = iota
	gbStateFEATURES
	gbStateSEQUENCE
	gbStateDONE
)

func ChannelFromGenbank(r io.Reader) <-chan Protein {
	out := make(chan Protein)
	go func() {
		defer close(out)
		state := gbStateSTART
		//		buf := bytes.Buffer{}
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			switch state {
			case gbStateSTART:
				if strings.HasPrefix(scanner.Text(), "FEATURES") {
					state = gbStateFEATURES
					continue
				}
				fmt.Println(scanner.Text())
			default:
				if strings.HasPrefix(scanner.Text(), "//") {
					state = gbStateSTART
					fmt.Println()
				}
			}
		}
		out <- Protein{
			Header:    "accession",
			AminoAcid: "VSATGGSSDALFAGVMEKVPSVANWIQVCG",
			Mass:      calculateMass("VSATGGSSDALFAGVMEKVPSVANWIQVCG"),
		}
	}()
	return out
}
