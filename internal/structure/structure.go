package structure

import (
	"bufio"
	"io"
	"strings"
	"strconv"
)

type Structure []Atom

type Atom struct {
	ID int
	TypeSymbol string
	Label Label
	PDBX string
	Cartesian Cartesian
	
}

type Label struct {
	AtomID string
	AltID  string
	CompID string
	AsymID string
	EntityID int
	SeqID int
}

type Cartesian struct {
	X float64
	Y float64
	Z float64
}

func NewStructure(r io.Reader) Structure {
	scanner := bufio.NewScanner(r)
	var s Structure
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ATOM") {
			fields := strings.Fields(scanner.Text())
			id, _ := strconv.Atoi(fields[1])
			eid, _ := strconv.Atoi(fields[7])
			sid, _ := strconv.Atoi(fields[8])
			x, _ := strconv.ParseFloat(fields[10], 64)
			y, _ := strconv.ParseFloat(fields[11], 64)
			z, _ := strconv.ParseFloat(fields[12], 64)
			a := Atom{
				ID: id,
				TypeSymbol: fields[2],
				Label: Label{
					AtomID: fields[3],
					AltID: fields[4],
					CompID: fields[5],
					AsymID: fields[6],
					EntityID: eid,
					SeqID: sid,
				},
				PDBX: fields[9],
				Cartesian: Cartesian{
					X: x,
					Y: y,
					Z: z,
				},
			}
			s = append(s, a)
		}
	}
	return s
}
