package protein

type Uniprot struct {
	EntryType        string `json:"entryType"`
	PrimaryAccession string `json:"primaryAccession"`
	UniProtkbID      string `json:"uniProtkbId"`
	EntryAudit       struct {
		FirstPublicDate          string `json:"firstPublicDate"`
		LastAnnotationUpdateDate string `json:"lastAnnotationUpdateDate"`
		LastSequenceUpdateDate   string `json:"lastSequenceUpdateDate"`
		EntryVersion             int    `json:"entryVersion"`
		SequenceVersion          int    `json:"sequenceVersion"`
	} `json:"entryAudit"`
	AnnotationScore float64 `json:"annotationScore"`
	Organism        struct {
		ScientificName string   `json:"scientificName"`
		CommonName     string   `json:"commonName"`
		Synonyms       []string `json:"synonyms"`
		TaxonID        int      `json:"taxonId"`
		Lineage        []string `json:"lineage"`
	} `json:"organism"`
	ProteinExistence   string `json:"proteinExistence"`
	ProteinDescription struct {
		RecommendedName struct {
			FullName struct {
				Evidences []struct {
					EvidenceCode string `json:"evidenceCode"`
					Source       string `json:"source"`
					ID           string `json:"id"`
				} `json:"evidences"`
				Value string `json:"value"`
			} `json:"fullName"`
			EcNumbers []struct {
				Evidences []struct {
					EvidenceCode string `json:"evidenceCode"`
					Source       string `json:"source"`
					ID           string `json:"id"`
				} `json:"evidences"`
				Value string `json:"value"`
			} `json:"ecNumbers"`
		} `json:"recommendedName"`
		AlternativeNames []struct {
			FullName struct {
				Evidences []struct {
					EvidenceCode string `json:"evidenceCode"`
					Source       string `json:"source"`
					ID           string `json:"id"`
				} `json:"evidences"`
				Value string `json:"value"`
			} `json:"fullName"`
			ShortNames []struct {
				Evidences []struct {
					EvidenceCode string `json:"evidenceCode"`
					Source       string `json:"source"`
					ID           string `json:"id"`
				} `json:"evidences"`
				Value string `json:"value"`
			} `json:"shortNames"`
		} `json:"alternativeNames"`
		Flag string `json:"flag"`
	} `json:"proteinDescription"`
	Comments []struct {
		Texts []struct {
			Evidences []struct {
				EvidenceCode string `json:"evidenceCode"`
				Source       string `json:"source"`
				ID           string `json:"id"`
			} `json:"evidences"`
			Value string `json:"value"`
		} `json:"texts,omitempty"`
		CommentType string `json:"commentType"`
		Reaction    struct {
			Name                    string `json:"name"`
			ReactionCrossReferences []struct {
				Database string `json:"database"`
				ID       string `json:"id"`
			} `json:"reactionCrossReferences"`
			Evidences []struct {
				EvidenceCode string `json:"evidenceCode"`
				Source       string `json:"source"`
				ID           string `json:"id"`
			} `json:"evidences"`
		} `json:"reaction,omitempty"`
		SubcellularLocations []struct {
			Location struct {
				Evidences []struct {
					EvidenceCode string `json:"evidenceCode"`
					Source       string `json:"source"`
					ID           string `json:"id"`
				} `json:"evidences"`
				Value string `json:"value"`
				ID    string `json:"id"`
			} `json:"location"`
		} `json:"subcellularLocations,omitempty"`
	} `json:"comments"`
	Features []struct {
		Type     string `json:"type"`
		Location struct {
			Start struct {
				Value    int    `json:"value"`
				Modifier string `json:"modifier"`
			} `json:"start"`
			End struct {
				Value    int    `json:"value"`
				Modifier string `json:"modifier"`
			} `json:"end"`
		} `json:"location"`
		Description string `json:"description"`
		Evidences   []struct {
			EvidenceCode string `json:"evidenceCode"`
		} `json:"evidences"`
		FeatureID           string `json:"featureId,omitempty"`
		AlternativeSequence struct {
			OriginalSequence     string   `json:"originalSequence"`
			AlternativeSequences []string `json:"alternativeSequences"`
		} `json:"alternativeSequence,omitempty"`
	} `json:"features"`
	Keywords []struct {
		ID       string `json:"id"`
		Category string `json:"category"`
		Name     string `json:"name"`
	} `json:"keywords"`
	References []struct {
		ReferenceNumber int `json:"referenceNumber"`
		Citation        struct {
			ID                      string   `json:"id"`
			CitationType            string   `json:"citationType"`
			Authors                 []string `json:"authors"`
			CitationCrossReferences []struct {
				Database string `json:"database"`
				ID       string `json:"id"`
			} `json:"citationCrossReferences"`
			Title           string `json:"title"`
			PublicationDate string `json:"publicationDate"`
			Journal         string `json:"journal"`
			FirstPage       string `json:"firstPage"`
			LastPage        string `json:"lastPage"`
			Volume          string `json:"volume"`
		} `json:"citation"`
		ReferencePositions []string `json:"referencePositions"`
		ReferenceComments  []struct {
			Evidences []struct {
				EvidenceCode string `json:"evidenceCode"`
				Source       string `json:"source"`
				ID           string `json:"id"`
			} `json:"evidences"`
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"referenceComments,omitempty"`
		Evidences []struct {
			EvidenceCode string `json:"evidenceCode"`
			Source       string `json:"source"`
			ID           string `json:"id"`
		} `json:"evidences"`
	} `json:"references"`
	UniProtKBCrossReferences []struct {
		Database   string `json:"database"`
		ID         string `json:"id"`
		Properties []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"properties"`
		Evidences []struct {
			EvidenceCode string `json:"evidenceCode"`
			Source       string `json:"source"`
			ID           string `json:"id"`
		} `json:"evidences,omitempty"`
	} `json:"uniProtKBCrossReferences"`
	Sequence struct {
		Value     string `json:"value"`
		Length    int    `json:"length"`
		MolWeight int    `json:"molWeight"`
		Crc64     string `json:"crc64"`
		Md5       string `json:"md5"`
	} `json:"sequence"`
	ExtraAttributes struct {
		CountByCommentType struct {
			Function            int `json:"FUNCTION"`
			CATALYTICACTIVITY   int `json:"CATALYTIC ACTIVITY"`
			ACTIVITYREGULATION  int `json:"ACTIVITY REGULATION"`
			SUBCELLULARLOCATION int `json:"SUBCELLULAR LOCATION"`
			Similarity          int `json:"SIMILARITY"`
		} `json:"countByCommentType"`
		CountByFeatureType struct {
			Signal        int `json:"Signal"`
			Chain         int `json:"Chain"`
			ActiveSite    int `json:"Active site"`
			Glycosylation int `json:"Glycosylation"`
			DisulfideBond int `json:"Disulfide bond"`
			Mutagenesis   int `json:"Mutagenesis"`
		} `json:"countByFeatureType"`
		UniParcID string `json:"uniParcId"`
	} `json:"extraAttributes"`
}
