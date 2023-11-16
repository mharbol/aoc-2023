package solution

import (
	"errors"

	"github.com/mharbol/aoc-2023/solution/days"
	"github.com/mharbol/aoc-2023/util"
)

var allSolutions = map[uint8]Solution{
    0: &days.Day00{},
}

type Solution interface {
	Day() uint8
	Part1([]string) (string, error)
	Part2([]string) (string, error)
}

func Solve(s Solution, part uint8) (string, error) {
	lines, err := util.ReadProblemInfo(s.Day())

	if err != nil {
		return "", err
	}

	switch part {
	case 1:
		return s.Part1(lines)
	case 2:
		return s.Part2(lines)
	default:
		return "", errors.New("Part must be 1 or 2")
	}
}
