package ex2

import (
	"strings"

	mr "github.com/vlad-a-barbu/mapreduce"
)

/*
	Mean number of palindrome words
*/

func Filter(word string) bool {
	var palindrome string
	for _, r := range word {
		palindrome = string(r) + palindrome
	}
	return strings.Compare(palindrome, word) == 0
}

func Solve(input [][]string) float64 {
	acc, _ := mr.MapReduce(input, Filter)
	return acc
}
