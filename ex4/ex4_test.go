package ex4

import "testing"

func TestSolve(t *testing.T) {
	input := [][]string{
		{"ana", "parc", "impare", "era", "copil"},
		{"cer", "program", "leu", "alee", "golang", "info"},
		{"inima", "impar", "apa", "eleve"}}

	got := Solve(input)
	want := 2.66

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
