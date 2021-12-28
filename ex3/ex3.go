package ex3

import (
	"strings"

	mr "github.com/vlad-a-barbu/mapreduce"
)

/*
	Mean number of words with:
		-> a 'p' after every vowel
*/

func Filter(word string) bool {
	check := false
	for _, r := range word {
		if check == true && r != 'p' {
			return false
		}
		if strings.ContainsRune("aeiouAEIOU", r) {
			check = true
		} else {
			check = false
		}
	}
	return true
}

func Solve(input [][]string) float64 {
	return mr.MapReduce(input, Filter)
}
