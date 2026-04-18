package dna

import "time"

// This type is not yet implemented.
type GenBank2 struct {
	Meta struct {
		Format struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			URL     string `json:"url"`
		} `json:"format"`
		Circular   bool      `json:"circular"`
		Locus      string    `json:"locus"`
		Length     int       `json:"length"`
		Datetime   time.Time `json:"datetime"`
		References []struct {
			Title   string `json:"title"`
			Index   int    `json:"index"`
			Range   []int  `json:"range"`
			Authors string `json:"authors"`
			Journal string `json:"journal"`
		} `json:"references"`
		Definition string `json:"definition"`
		Accession  string `json:"accession"`
		Version    string `json:"version"`
		Keywords   string `json:"keywords"`
		Source     struct {
			Range        []int  `json:"range"`
			Strain       string `json:"strain"`
			DbXref       string `json:"db_xref"`
			Organism     string `json:"organism"`
			IdentifiedBy string `json:"identified_by"`
			MolType      string `json:"mol_type"`
		} `json:"source"`
		Organism string `json:"organism"`
	} `json:"meta"`
	Features []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Range       []int  `json:"range"`
		Product     string `json:"product"`
		Note        string `json:"note,omitempty"`
		CodonStart  string `json:"codon_start,omitempty"`
		ProteinID   string `json:"protein_id,omitempty"`
		Function    string `json:"function,omitempty"`
		Translation string `json:"translation,omitempty"`
	} `json:"features"`
	Origin string `json:"origin"`
}
