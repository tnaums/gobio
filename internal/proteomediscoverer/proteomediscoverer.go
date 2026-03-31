package proteomediscoverer

import (
	"encoding/csv"
	"fmt"
	"io"
	//"log"
	"strings"
)

type ProteomeDiscoverer struct {
	Accession string
}

func GetAccession(entry string) (ProteomeDiscoverer, error) {
	columns := strings.Split(entry, ",")
	fmt.Println(len(columns))
	return ProteomeDiscoverer{
		Accession: columns[3],
	}, nil
}

func ParseCSV(f io.Reader) ([]ProteomeDiscoverer, error) {
	r := csv.NewReader(f)
	var records []ProteomeDiscoverer
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return records, err
		}
		if record[0] == "Checked"{
			continue
		}

		pd := ProteomeDiscoverer{
			Accession: record[3],
		}
		records = append(records, pd)
	}
	return records, nil
}
