package solution

import "github.com/mharbol/aoc-2023/solution/days"

type Solution interface {
	Day() uint8
	Part1() (string, error)
	Part2() (string, error)
}

var allSolutions = map[uint8]Solution {
    0: &days.Day00{},
}

type expectedPair struct {
	part1 string
	part2 string
}

var allExpected = map[uint8]expectedPair{
    0: {"", ""},
}
