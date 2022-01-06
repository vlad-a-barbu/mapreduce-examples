package ex5

import "testing"

func TestSolve(t *testing.T) {
	input := [][]string{
		{"AcasA", "CasA", "FacultatE", "SisTemE", "distribuite"},
		{"GolanG", "map", "reduce", "Problema", "TemA", "ProieCt"},
		{"LicentA", "semestru", "ALGORitM", "StuDent"}}

	got := Solve(input)
	want := 1.66

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
