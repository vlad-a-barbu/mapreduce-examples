package ex5

import (
	"unicode"

	mr "github.com/vlad-a-barbu/mapreduce"
)

/*
	Mean number of words that:
		- begin and end with an uppercase letter
		- even number of lowercase letter
*/

func Filter(word string) bool {
	runes := []rune(word)

	cnt := 0
	for _, r := range runes {
		if unicode.IsLower(r) {
			cnt++
		}
	}

	return unicode.IsUpper(runes[0]) &&
		unicode.IsUpper(runes[len(word)-1]) &&
		cnt%2 == 0
}

func Solve(input [][]string) float64 {
	acc, _ := mr.MapReduce(input, Filter)
	return acc
}
