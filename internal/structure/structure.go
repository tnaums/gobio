package structure

import (
	"bufio"
	//	"fmt"
	"io"
	"strconv"
	"strings"
)

type Structure []Residue

type Atom struct {
	ID         int
	TypeSymbol string
	Label      Label
	Cartesian  Cartesian
	Occupancy  float64
	B          float64
	Author     Author
	PDBX       PDBX
}

type PDBX struct {
	InsCode     string
	PDBModelNum int
}

type Label struct {
	AtomID   string
	AltID    string
	CompID   string
	AsymID   string
	EntityID int
	SeqID    int
}

type Cartesian struct {
	X float64
	Y float64
	Z float64
}

type Author struct {
	SeqID  int
	AsymID string
}

type Chain struct {
	Sequence []Residue
}

type Residue struct {
	Chain    string
	Position int
	Atoms    []Atom
}

func NewStructure(r io.Reader) Structure {
	scanner := bufio.NewScanner(r)
	var structure Structure
	var aminoAcid Residue
	var currentResidue int
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ATOM") {
			atom := NewAtom(scanner.Text())
			if currentResidue != atom.Label.SeqID {
				structure = append(structure, aminoAcid)
				aminoAcid = Residue{
					Chain:    atom.Label.AsymID,
					Position: atom.Label.SeqID,
					Atoms:    make([]Atom, 0),
				}
				currentResidue = atom.Label.SeqID
			}
			aminoAcid.Atoms = append(aminoAcid.Atoms, atom)
		}
	}
	structure = append(structure, aminoAcid)
	return structure
}

func NewAtom(entry string) Atom {
	fields := strings.Fields(entry)
	id, _ := strconv.Atoi(fields[1])
	eid, _ := strconv.Atoi(fields[7])
	sid, _ := strconv.Atoi(fields[8])
	x, _ := strconv.ParseFloat(fields[10], 64)
	y, _ := strconv.ParseFloat(fields[11], 64)
	z, _ := strconv.ParseFloat(fields[12], 64)
	o, _ := strconv.ParseFloat(fields[13], 64)
	b, _ := strconv.ParseFloat(fields[14], 64)
	sid2, _ := strconv.Atoi(fields[15])
	pdbx, _ := strconv.Atoi(fields[17])
	a := Atom{
		ID:         id,
		TypeSymbol: fields[2],
		Label: Label{
			AtomID:   fields[3],
			AltID:    fields[4],
			CompID:   fields[5],
			AsymID:   fields[6],
			EntityID: eid,
			SeqID:    sid,
		},
		Cartesian: Cartesian{
			X: x,
			Y: y,
			Z: z,
		},
		Occupancy: o,
		B:         b,
		Author: Author{
			SeqID:  sid2,
			AsymID: fields[16],
		},
		PDBX: PDBX{
			InsCode:     fields[9],
			PDBModelNum: pdbx,
		},
	}
	return a
}
