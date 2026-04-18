package dna

import (
	"bufio"
	"bytes"
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
	state    genBankState
}

func NewGenBank(r io.Reader) GenBank {
	g := GenBank{
		state: genbankStateNone,
	}
	builder := strings.Builder{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ORIGIN") {
			g.state = genbankStateOrigin
			continue
		}
		if strings.HasPrefix(scanner.Text(), "//") {
			g.state = genbankStateDone
			sequence := strings.Join(strings.Fields(builder.String()), "")
			g.Sequence = NewDNAFromSequence("genbank", sequence)
		}
		if g.state == genbankStateOrigin {
			trimmed := bytes.Trim(scanner.Bytes(), " 0123456789")
			builder.Write(trimmed)
		}
	}
	return g
}
