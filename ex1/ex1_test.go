package ex1

import "testing"

func TestSolve(t *testing.T) {
	input := [][]string{
		{"aabbb", "ebep", "blablablaa", "hijk", "wsww"},
		{"abba", "eeeppp", "cocor", "ppppppaa", "qwerty", "acasq"},
		{"lalala", "lalal", "papapa", "papap"}}

	got := Solve(input)
	want := 2.33

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
