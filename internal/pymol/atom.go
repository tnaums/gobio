package pymol

import (
	"strconv"
	"strings"
)

// Atom holds information parsed from ATOM line in cif file
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

// Label portion of atom
type Label struct {
	AtomID   string
	AltID    string
	CompID   string
	AsymID   string
	EntityID int
	SeqID    int
}

// Cartesian portion of atom
type Cartesian struct {
	X float64
	Y float64
	Z float64
}

// Author portion of atom
type Author struct {
	SeqID  int
	AsymID string
}

// PDBX portion of atom
type PDBX struct {
	InsCode     string
	PDBModelNum int
}

// NewAtom parses information from an ATOM line in a cif
// protein structure file and returns an Atom struct.
func NewAtom(entry string) Atom {
	// split line into 17 values
	fields := strings.Fields(entry)
	// convert integer and float values from strings
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
	// create Atom entry
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
