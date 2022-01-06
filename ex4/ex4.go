package ex4

import (
	"strings"

	mr "github.com/vlad-a-barbu/mapreduce"
)

/*
	Mean number of words that begin and end with a vowel
*/

func IsVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

func Filter(word string) bool {
	runes := []rune(word)
	return IsVowel(runes[0]) && IsVowel(runes[len(word)-1])
}

func Solve(input [][]string) float64 {
	acc, _ := mr.MapReduce(input, Filter)
	return acc
}
