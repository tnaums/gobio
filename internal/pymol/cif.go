package pymol

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Map that converts the 3 letter amino acid codes found in
// ATOM records to single amino acid codes. Used to create
// fasta protein sequences from cif files.
var ThreeToOne = map[string]byte{
	"ALA": 'A', "LEU": 'L',
	"ARG": 'R', "LYS": 'K',
	"ASN": 'N', "MET": 'M',
	"ASP": 'D', "PHE": 'F',
	"CYS": 'C', "PRO": 'P',
	"GLN": 'Q', "SER": 'S',
	"GLU": 'E', "THR": 'T',
	"GLY": 'G', "TRP": 'W',
	"HIS": 'H', "TYR": 'Y',
	"ILE": 'I', "VAL": 'V',
}

// Function that creates protein fasta files for each chain in a cif
// file. The returned *bytes.Buffer can be passed to
// protein.ProteinChannelFasta as the io.Reader.
func SequenceFromCIF(r io.Reader) *bytes.Buffer {
	buf := bytes.Buffer{}
	scanner := bufio.NewScanner(r)
	currentChain := ""
	currentAA := "0"
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ATOM") {
			fields := strings.Fields(scanner.Text())
			if fields[6] != currentChain {
				buf.WriteString(fmt.Sprintf("\n>%s\n", fields[6]))
				currentChain = fields[6]
			}
			if fields[8] != currentAA {
				buf.WriteByte(ThreeToOne[fields[5]])
				currentAA = fields[8]
			}
		}
	}
	return &buf
}

// Residue contains information for an amino acid.
type Residue struct {
	AminoAcid string
	Position  int
	IDStart   int
	IDEnd     int
}

// ChainMap keys are sequence number for an amino acid in a chain.
// Values are Residue struct for that amino acid. Used to convert
// amino acid numbers to atom id numbers.
type ChainMap map[int]Residue

// Create a ChainMap from the ATOM field of a cif file.
func NewChainMap(r io.Reader, chain string) ChainMap{
	scanner := bufio.NewScanner(r)
	currentResidue := 0
	id := 0
	var residue Residue
	m := make(map[int]Residue)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ATOM") {
			fields := strings.Fields(scanner.Text())
			if fields[6] == chain {
				seqid, _ := strconv.Atoi(fields[8])
				id, _ = strconv.Atoi(fields[1])
				// if it is the first one
				if currentResidue == 0 {
					currentResidue += 1
					residue = Residue{
						AminoAcid: fields[5],
						Position:  currentResidue,
						IDStart:   id,
					}
				}
				if seqid != currentResidue {
					residue.IDEnd = id - 1
					m[currentResidue] = residue
					currentResidue += 1
					residue = Residue{
						AminoAcid: fields[5],
						Position:  currentResidue,
						IDStart:   id,
					}
				}
			}
		}
	}
	residue.IDEnd = id
	m[currentResidue] = residue
	return m
}

// Keys are atom id. Values are the Atom struct containing all 17 fields of
// information parsed from ATOM lines of cif file.
type Structure map[int]Atom

// Creates a new Structure map from a cif file.
func NewStructure(r io.Reader) Structure {
	scanner := bufio.NewScanner(r)
	structure := make(Structure, 0)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ATOM") {
			atom := NewAtom(scanner.Text())
			structure[atom.ID] = atom
		}
	}
	return structure
}
