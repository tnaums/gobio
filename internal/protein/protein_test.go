package protein

import (
	"fmt"
	"testing"
)

func TestNewProteinFromFasta(t *testing.T) {
	fastaFile := "small.fa"
	expectedName := "jgi|Fusve2|1|FVEG_14560T0"
	expectedSequence := "MRWPFLDLSATALLSTASHAFSETRGSQRYAASTQCYILYMIGYVEGHWTT"
	gotProtSlice, _ := NewProteinFromFasta(fastaFile)
	if len(gotProtSlice) != 1 {
		t.Errorf("expected %d sequences, got %d", 1, len(gotProtSlice))
	}
	if gotProtSlice[0].Header != expectedName {
		t.Errorf("expected header name '%s', got '%s'", expectedName, gotProtSlice[0].Header)
	}
	if gotProtSlice[0].AminoAcid != expectedSequence {
		t.Errorf("\nexpected sequence %s\n     got %s\n", expectedSequence, gotProtSlice[0].AminoAcid)
	}
}

func ExampleProteinPipeFasta() {
	fmt.Println("This is an example of ProteinPipeFasta")
	// Output: This is an example of ProteinPipeFasta
}
