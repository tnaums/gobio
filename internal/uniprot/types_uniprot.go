package uniprot

import (
	"fmt"

	"github.com/tnaums/gobio/internal/protein"
)

type UniprotAPI struct {
	ResponseType string
	IdList       string
}

type UniprotComplete struct {
	JSON     UniprotRecord
	Flatfile []byte
}

func (u UniprotComplete) GetFasta() protein.Protein {
	header := fmt.Sprintf("%s|%s|%s", u.JSON.Accession, u.JSON.Organism.Names[0].Value, u.JSON.Protein.RecommendedName.FullName.Value)
	protein := protein.NewProtein(header, u.JSON.Sequence.Sequence)
	return protein
}

func (u UniprotComplete) GetFlatFile() []byte {
	return u.Flatfile
}

func (u UniprotComplete) PrintFeatures() {
	for _, feature := range u.JSON.Features {
		fmt.Printf("       Type: %s\n", feature.Type)
		fmt.Printf("   Category: %s\n", feature.Category)
		fmt.Printf("Description: %s\n", feature.Description)
		fmt.Printf("      Begin: %s\n", feature.Begin)
		fmt.Printf("        End: %s\n", feature.End)
		fmt.Printf("   Molecule:  %s\n", feature.Molecule)
		fmt.Printf("  Evidences: %s\n", feature.Evidences)
		fmt.Println()
	}

	return
}

type UniprotRecord struct {
	Accession        string `json:"accession"`
	ID               string `json:"id"`
	ProteinExistence string `json:"proteinExistence"`
	Info             struct {
		Type     string `json:"type"`
		Created  string `json:"created"`
		Modified string `json:"modified"`
		Version  int    `json:"version"`
	} `json:"info"`
	Organism struct {
		Taxonomy int `json:"taxonomy"`
		Names    []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"names"`
		Lineage []string `json:"lineage"`
	} `json:"organism"`
	Protein struct {
		RecommendedName struct {
			FullName struct {
				Value     string `json:"value"`
				Evidences []struct {
					Code   string `json:"code"`
					Source struct {
						Name           string `json:"name"`
						ID             string `json:"id"`
						URL            string `json:"url"`
						AlternativeURL string `json:"alternativeUrl"`
					} `json:"source"`
				} `json:"evidences"`
			} `json:"fullName"`
			EcNumber []struct {
				Value     string `json:"value"`
				Evidences []struct {
					Code   string `json:"code"`
					Source struct {
						Name           string `json:"name"`
						ID             string `json:"id"`
						URL            string `json:"url"`
						AlternativeURL string `json:"alternativeUrl"`
					} `json:"source"`
				} `json:"evidences"`
			} `json:"ecNumber"`
		} `json:"recommendedName"`
		AlternativeName []struct {
			FullName struct {
				Value     string `json:"value"`
				Evidences []struct {
					Code   string `json:"code"`
					Source struct {
						Name           string `json:"name"`
						ID             string `json:"id"`
						URL            string `json:"url"`
						AlternativeURL string `json:"alternativeUrl"`
					} `json:"source"`
				} `json:"evidences"`
			} `json:"fullName"`
			ShortName []struct {
				Value     string `json:"value"`
				Evidences []struct {
					Code   string `json:"code"`
					Source struct {
						Name           string `json:"name"`
						ID             string `json:"id"`
						URL            string `json:"url"`
						AlternativeURL string `json:"alternativeUrl"`
					} `json:"source"`
				} `json:"evidences"`
			} `json:"shortName"`
		} `json:"alternativeName"`
	} `json:"protein"`
	Comments []struct {
		Type string `json:"type"`
		Text []struct {
			Value     string `json:"value"`
			Evidences []struct {
				Code   string `json:"code"`
				Source struct {
					Name           string `json:"name"`
					ID             string `json:"id"`
					URL            string `json:"url"`
					AlternativeURL string `json:"alternativeUrl"`
				} `json:"source"`
			} `json:"evidences"`
		} `json:"text,omitempty"`
		Reaction struct {
			Name         string `json:"name"`
			DbReferences []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"dbReferences"`
			Evidences []struct {
				Code   string `json:"code"`
				Source struct {
					Name           string `json:"name"`
					ID             string `json:"id"`
					URL            string `json:"url"`
					AlternativeURL string `json:"alternativeUrl"`
				} `json:"source"`
			} `json:"evidences"`
		} `json:"reaction,omitempty"`
		Locations []struct {
			Location struct {
				Value     string `json:"value"`
				Evidences []struct {
					Code   string `json:"code"`
					Source struct {
						Name           string `json:"name"`
						ID             string `json:"id"`
						URL            string `json:"url"`
						AlternativeURL string `json:"alternativeUrl"`
					} `json:"source"`
				} `json:"evidences"`
			} `json:"location"`
		} `json:"locations,omitempty"`
	} `json:"comments"`
	Features []struct {
		Type        string `json:"type"`
		Category    string `json:"category"`
		Description string `json:"description"`
		Begin       string `json:"begin"`
		End         string `json:"end"`
		Molecule    string `json:"molecule"`
		Evidences   []struct {
			Code string `json:"code"`
		} `json:"evidences"`
		FtID                string `json:"ftId,omitempty"`
		AlternativeSequence string `json:"alternativeSequence,omitempty"`
	} `json:"features"`
	DbReferences []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Properties struct {
			MoleculeType      string `json:"molecule type"`
			ProteinSequenceID string `json:"protein sequence ID"`
		} `json:"properties,omitempty"`
		Properties0 struct {
			Term   string `json:"term"`
			Source string `json:"source"`
		} `json:"properties,omitempty"`
		Evidences []struct {
			Code   string `json:"code"`
			Source struct {
				Name           string `json:"name"`
				ID             string `json:"id"`
				URL            string `json:"url"`
				AlternativeURL string `json:"alternativeUrl"`
			} `json:"source"`
		} `json:"evidences,omitempty"`
		Properties1 struct {
			Term   string `json:"term"`
			Source string `json:"source"`
		} `json:"properties,omitempty"`
		Properties2 struct {
			Term   string `json:"term"`
			Source string `json:"source"`
		} `json:"properties,omitempty"`
		Properties3 struct {
			MatchStatus string `json:"match status"`
			EntryName   string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties4 struct {
			EntryName string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties5 struct {
			EntryName string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties6 struct {
			EntryName string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties7 struct {
			EntryName string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties8 struct {
			MatchStatus string `json:"match status"`
			EntryName   string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties9 struct {
			MatchStatus string `json:"match status"`
			EntryName   string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties10 struct {
			MatchStatus string `json:"match status"`
			EntryName   string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties11 struct {
			MatchStatus string `json:"match status"`
			EntryName   string `json:"entry name"`
		} `json:"properties,omitempty"`
		Properties12 struct {
			MatchStatus string `json:"match status"`
			EntryName   string `json:"entry name"`
		} `json:"properties,omitempty"`
	} `json:"dbReferences"`
	Keywords []struct {
		Value string `json:"value"`
	} `json:"keywords"`
	References []struct {
		Citation struct {
			Type            string   `json:"type"`
			PublicationDate string   `json:"publicationDate"`
			Title           string   `json:"title"`
			Authors         []string `json:"authors"`
			Publication     struct {
				JournalName string `json:"journalName"`
			} `json:"publication"`
			Location struct {
				Volume    string `json:"volume"`
				FirstPage string `json:"firstPage"`
				LastPage  string `json:"lastPage"`
			} `json:"location"`
			DbReferences []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"dbReferences"`
		} `json:"citation"`
		Source struct {
			Strain []struct {
				Value     string `json:"value"`
				Evidences []struct {
					Code   string `json:"code"`
					Source struct {
						Name           string `json:"name"`
						ID             string `json:"id"`
						URL            string `json:"url"`
						AlternativeURL string `json:"alternativeUrl"`
					} `json:"source"`
				} `json:"evidences"`
			} `json:"strain"`
		} `json:"source,omitempty"`
		Scope     []string `json:"scope"`
		Evidences []struct {
			Code   string `json:"code"`
			Source struct {
				Name string `json:"name"`
				ID   string `json:"id"`
				URL  string `json:"url"`
			} `json:"source"`
		} `json:"evidences"`
	} `json:"references"`
	Sequence struct {
		Version  int    `json:"version"`
		Length   int    `json:"length"`
		Mass     int    `json:"mass"`
		Modified string `json:"modified"`
		Sequence string `json:"sequence"`
	} `json:"sequence"`
}
