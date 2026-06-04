package dna

import (
	"bufio"
	"bytes"
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

func ChannelFromGenbank(r io.Reader) <-chan DNA {
	out := make(chan DNA)
	go func() {
		defer close(out)
		state := gbStateSTART
		buf := bytes.Buffer{}
		accession := ""
		definition := ""
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			switch state {
			case gbStateSTART:
				if strings.HasPrefix(scanner.Text(), "FEATURES") {
					state = gbStateFEATURES
					continue
				}
				if strings.HasPrefix(scanner.Text(), "ACCESSION") {
					a := strings.Fields(scanner.Text())
					accession = a[1]
					continue
				}
				if strings.HasPrefix(scanner.Text(), "DEFINITION") {
					definition = strings.TrimSpace(strings.TrimPrefix(scanner.Text(), "DEFINITION"))
					definition = strings.TrimSuffix(definition, ".")
					continue
				}
			case gbStateFEATURES:
				if strings.HasPrefix(scanner.Text(), "ORIGIN") {
					state = gbStateSEQUENCE
					continue
				}
				trimmed := strings.Trim(scanner.Text(), " ")
				if strings.HasPrefix(trimmed, "CDS") {
					fmt.Printf("Found a CDS: %s\n", trimmed)
				}
			case gbStateSEQUENCE:
				if strings.HasPrefix(scanner.Text(), "//") {
					sequence := strings.ToUpper(strings.Join(strings.Fields(buf.String()), ""))
					newDNA := DNA{
						Header: fmt.Sprintf("%s|%s", accession, definition),
						Parent: sequence,
						Complement: reverseComplement(sequence),
					}
					newDNA.Orfs = newDNA.Translate()
					out <- newDNA
					state = gbStateSTART
					sequence = ""
					accession = ""
					definition = ""
					buf.Truncate(0)
					continue
				}
				trimmed := bytes.Trim(scanner.Bytes(), " 0123456789")
				buf.Write(trimmed)
				continue
			}
		}
	}()
	return out
}
