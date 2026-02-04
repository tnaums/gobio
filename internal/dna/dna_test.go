package dna

import (
	//	"fmt"
	"testing"
)

func TestFastaParser(t *testing.T) {
	name, sequence := FastaParser("test_file.fa")
	ename := "pTest"
	esequence := "CTCGAGCTTAATTAACAACACCATTTGTCGAGAAATCATAAAAAATTTATTTGCTTTGTGAGCGGATAACAATTAT"
	if name != ename {
		t.Errorf("expected %s but got %s\n", ename, name)
	}
	if sequence != esequence {
		t.Errorf("\nexpected: %s\n     got: %s\n", esequence, sequence)
	}
}


func TestReverse(t *testing.T) {
	t.Run("reversing a simple string", func(t *testing.T) {
		got := reverse("My name is Todd")
		want := "ddoT si eman yM"
		assertCorrectMessage(t, got, want)
	})
	t.Run("trying with unicode", func(t *testing.T) {
		got := reverse("Hello, 世界")
		want := "界世 ,olleH"
		assertCorrectMessage(t, got, want)
	})
}

func TestReverseComplement(t *testing.T) {
	sequence := "CTCGAGCTTAATTAACAACACCATTTGTCGAGAAATCATAAAAAATTTATTTGCTTTGTGAGCGGATAACAATTAT"
	got := ReverseComplement(sequence)
	want := "ATAATTGTTATCCGCTCACAAAGCAAATAAATTTTTTATGATTTCTCGACAAATGGTGTTGTTAATTAAGCTCGAG"
	if got != want {
		t.Errorf("\nexpected %s\n     got %s\n", want, got)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewDnaFromSequence(t *testing.T) {
	sequence := "CTCGAGCTTAATTAACAACACCATTTGTCGAGAAATCATAAAAAATTTATTTGCTTTGTGAGCGGATAACAATTAT"
	got := NewDnaFromSequence(sequence)
	wantP := "CTCGAGCTTAATTAACAACACCATTTGTCGAGAAATCATAAAAAATTTATTTGCTTTGTGAGCGGATAACAATTAT"
	wantC := "ATAATTGTTATCCGCTCACAAAGCAAATAAATTTTTTATGATTTCTCGACAAATGGTGTTGTTAATTAAGCTCGAG"
	if got.Parent != wantP {
		t.Errorf("\nexpected %s\n     got %s\n", wantP, got.Parent)
	}
	if got.Complement != wantC {
		t.Errorf("\nexpected %s\n     got %s\n", wantC, got.Complement)
	}
}

func TestNewDnaFromFasta(t *testing.T) {
	fastaFile := "test_file.fa"
	expectedName := "pTest"
	expectedSequence := "CTCGAGCTTAATTAACAACACCATTTGTCGAGAAATCATAAAAAATTTATTTGCTTTGTGAGCGGATAACAATTAT"
	gotDna := NewDnaFromFasta(fastaFile)
	if gotDna.File != fastaFile {
		t.Errorf("expected filename %s, got %s", fastaFile, gotDna.File)
	}
	if gotDna.Name != expectedName {
		t.Errorf("expected header name '%s', got '%s'", expectedName, gotDna.Name)
	}
	if gotDna.Parent != expectedSequence {
		t.Errorf("\nexpected %s\n     got %s\n", expectedSequence, gotDna.Parent)
	}
}
