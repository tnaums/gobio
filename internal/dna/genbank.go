package dna

import (
	"bufio"
	"bytes"
	//	"fmt"
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
	state    genBankState
}

func NewGenBank(r io.Reader) GenBank {
	g := GenBank{
		state: genbankStateNone,
	}
	buf := bytes.Buffer{}
	features := bytes.Buffer{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ORIGIN") {
			g.state = genbankStateOrigin
			g.Features = features.Bytes()
			continue
		}
		if strings.HasPrefix(scanner.Text(), "//") {
			g.state = genbankStateDone
			sequence := strings.Join(strings.Fields(buf.String()), "")
			g.Sequence = NewDNAFromSequence("genbank", sequence)
		}
		if strings.HasPrefix(scanner.Text(), "FEATURES") {
			g.state = genbankStateFeatures
		}
		if g.state == genbankStateOrigin {
			trimmed := bytes.Trim(scanner.Bytes(), " 0123456789")
			buf.Write(trimmed)
		}
		if g.state == genbankStateFeatures {
			features.Write(scanner.Bytes())
			features.Write([]byte("\n"))
		}
	}
	return g
}
