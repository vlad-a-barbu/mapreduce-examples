package ex2

import "testing"

func TestSolve(t *testing.T) {
	input := [][]string{
		{"a1551a", "parc", "ana", "minim", "1pcl3"},
		{"calabalac", "tivit", "leu", "zece10", "ploaie", "9ana9"},
		{"lalalal", "tema", "papa", "ger"}}

	got := Solve(input)
	want := 2.33

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
