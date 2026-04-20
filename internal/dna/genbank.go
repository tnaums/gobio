package dna

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type genBankState int

const (
	genbankStateNone genBankState = iota
	genbankStateLocus
	genbankStateDefinition
	genbankStateAccession
	genbankStateVersion
	genbankStateKeywords
	genbankStateSource
	genbankStateReference
	genbankStateFeatures
	genbankStateOrigin
	genbankStateDone
)

type GenBank struct {
	Sequence DNA
	Features []byte
	Accession string
	Definition string
	state    genBankState
}

// Parses a GenBank file containing a single dna sequence and returns a GenBank struct.
func NewGenBank(r io.Reader) GenBank {
	g := GenBank{
		state: genbankStateNone,
	}
	buf := bytes.Buffer{}
	features := bytes.Buffer{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "//") {
			g.state = genbankStateDone
			sequence := strings.Join(strings.Fields(buf.String()), "")
			header := fmt.Sprintf("%s|%dbp|%s", g.Accession,  len(sequence), g.Definition)
			g.Sequence = NewDNAFromSequence(header, sequence)
		}		
		if strings.HasPrefix(scanner.Text(), "ORIGIN") {
			g.state = genbankStateOrigin
			g.Features = features.Bytes()
			continue
		}
		if strings.HasPrefix(scanner.Text(), "FEATURES") {
			g.state = genbankStateFeatures
			continue
		}
		if strings.HasPrefix(scanner.Text(), "ACCESSION") {
			a := strings.Fields(scanner.Text())
			g.Accession = a[1]
			g.state = genbankStateAccession
			continue
		}
		if strings.HasPrefix(scanner.Text(), "DEFINITION") {
			d := strings.TrimSpace(strings.TrimPrefix(scanner.Text(), "DEFINITION"))
			g.Definition = d
			g.state = genbankStateDefinition
			continue
		}		
		if g.state == genbankStateOrigin {
			trimmed := bytes.Trim(scanner.Bytes(), " 0123456789")
			buf.Write(trimmed)
			continue
		}
		if g.state == genbankStateFeatures {
			features.Write(scanner.Bytes())
			features.Write([]byte("\n"))
			continue
		}
	}
	return g
}
