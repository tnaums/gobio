package alphafold

type alphafoldSummary []struct {
	ToolUsed                   string    `json:"toolUsed"`
	ProviderID                 string    `json:"providerId"`
	EntityType                 string    `json:"entityType"`
	IsUniProt                  bool      `json:"isUniProt"`
	ModelEntityID              string    `json:"modelEntityId"`
	ModelCreatedDate           time.Time `json:"modelCreatedDate"`
	SequenceVersionDate        time.Time `json:"sequenceVersionDate"`
	GlobalMetricValue          float64   `json:"globalMetricValue"`
	FractionPlddtVeryLow       float64   `json:"fractionPlddtVeryLow"`
	FractionPlddtLow           float64   `json:"fractionPlddtLow"`
	FractionPlddtConfident     float64   `json:"fractionPlddtConfident"`
	FractionPlddtVeryHigh      float64   `json:"fractionPlddtVeryHigh"`
	LatestVersion              int       `json:"latestVersion"`
	AllVersions                []int     `json:"allVersions"`
	Sequence                   string    `json:"sequence"`
	SequenceStart              int       `json:"sequenceStart"`
	SequenceEnd                int       `json:"sequenceEnd"`
	SequenceChecksum           string    `json:"sequenceChecksum"`
	IsUniProtReviewed          bool      `json:"isUniProtReviewed"`
	Gene                       string    `json:"gene"`
	UniprotAccession           string    `json:"uniprotAccession"`
	UniprotID                  string    `json:"uniprotId"`
	UniprotDescription         string    `json:"uniprotDescription"`
	TaxID                      int       `json:"taxId"`
	OrganismScientificName     string    `json:"organismScientificName"`
	IsUniProtReferenceProteome bool      `json:"isUniProtReferenceProteome"`
	BcifURL                    string    `json:"bcifUrl"`
	CifURL                     string    `json:"cifUrl"`
	PdbURL                     string    `json:"pdbUrl"`
	PaeImageURL                string    `json:"paeImageUrl"`
	MsaURL                     string    `json:"msaUrl"`
	PlddtDocURL                string    `json:"plddtDocUrl"`
	PaeDocURL                  string    `json:"paeDocUrl"`
	IsComplex                  bool      `json:"isComplex"`
	EntryID                    string    `json:"entryId"`
	UniprotSequence            string    `json:"uniprotSequence"`
	UniprotStart               int       `json:"uniprotStart"`
	UniprotEnd                 int       `json:"uniprotEnd"`
	IsReferenceProteome        bool      `json:"isReferenceProteome"`
	IsReviewed                 bool      `json:"isReviewed"`
}
