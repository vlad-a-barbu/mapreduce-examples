package ex1

import (
	"strings"

	mr "github.com/vlad-a-barbu/mapreduce"
)

/*
	Mean number of words with:
		-> even number of vowels
		-> number of consonants is a multiple of 3
*/

func Filter(word string) bool {

	vowels := 0
	consonants := 0

	for _, r := range word {
		if strings.ContainsRune("aeiouAEIOU", r) {
			vowels++
		} else {
			consonants++
		}
	}

	if vowels%2 == 0 && consonants%3 == 0 {
		return true
	}

	return false
}

func Solve(input [][]string) float64 {
	acc, _ := mr.MapReduce(input, Filter)
	return acc
}
