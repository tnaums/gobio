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

type fState int

const (
	fStateNONE gbState = iota
	fStateCDS
	fStateTRANSLATION
)

func ChannelFromGenbank(r io.Reader) <-chan DNA {
	out := make(chan DNA)
	go func() {
		defer close(out)
		state := gbStateSTART
		fstate := fStateNONE
		buf := bytes.Buffer{}
		cds := bytes.Buffer{}
		orfs := []Orf{}
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
				// Finished with FEATURES and move on to SEQUENCE
				if strings.HasPrefix(scanner.Text(), "ORIGIN") {
					state = gbStateSEQUENCE
					fstate = fStateNONE

					continue
				}

				trimmed := strings.Trim(scanner.Text(), " ")
				switch fstate {
				case fStateNONE:
					if strings.HasPrefix(trimmed, "CDS") {
						fstate = fStateCDS
					}
					continue
				case fStateCDS:
					if strings.HasPrefix(trimmed, "/translation=") {
						sequence := strings.TrimPrefix(trimmed, "/translation=\"")
						cds.Write([]byte(sequence))
						fstate = fStateTRANSLATION
					}
					continue
				case fStateTRANSLATION:
					if strings.Index(trimmed, "\"") != -1 {
						cds.Write([]byte(trimmed[:len(trimmed)-1]))
						newOrf := Orf{
							Strand:    "strand",
							Frame:     0,
							AminoAcid: cds.String(),
						}
						orfs = append(orfs, newOrf)
						fstate = fStateNONE
						cds.Truncate(0)
						continue
					}
					cds.Write([]byte(trimmed))
				}
			case gbStateSEQUENCE:
				if strings.HasPrefix(scanner.Text(), "//") {
					sequence := strings.ToUpper(strings.Join(strings.Fields(buf.String()), ""))
					out <- DNA{
						Header:     fmt.Sprintf("%s|%s", accession, definition),
						Parent:     sequence,
						Complement: reverseComplement(sequence),
						Orfs: orfs,
					}

					state = gbStateSTART
					fstate = fStateNONE
					sequence = ""
					accession = ""
					definition = ""
					orfs = []Orf{}
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
