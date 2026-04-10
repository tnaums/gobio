package structure2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

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

func SequenceFromCIF(r io.Reader) string {
	builder := strings.Builder{}
	scanner := bufio.NewScanner(r)
	currentChain := ""
	currentAA := "0"
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "ATOM") {
			fields := strings.Fields(scanner.Text())
			if fields[6] != currentChain {
				builder.WriteString(fmt.Sprintf("\n>%s\n", fields[6]))
				currentChain = fields[6]
			}
			if fields[8] != currentAA {
				builder.WriteByte(ThreeToOne[fields[5]])
				currentAA = fields[8]
			}
		}
	}
	return builder.String()
}
