package ex3

import "testing"

func TestSolve(t *testing.T) {
	input := [][]string{
		{"apap", "paprc", "apnap", "mipnipm", "copil"},
		{"cepr", "program", "lepu", "zepcep", "golang", "tema"},
		{"par", "impar", "papap", "gepr"}}

	got := Solve(input)
	want := 3.0

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
