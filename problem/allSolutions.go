package problem

import "github.com/mharbol/aoc-2023/solution"

var allSolutions map[uint8]solution.Solution = nil

func AllSolutions() *map[uint8]solution.Solution {
	if allSolutions == nil {
		allSolutions = makeAllSolutions()
	}
	return &allSolutions
}

func makeAllSolutions() map[uint8]solution.Solution {

	allSolutions = make(map[uint8]solution.Solution)

	allSolutions[0] = &Day00{}

	return allSolutions
}
