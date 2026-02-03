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

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
