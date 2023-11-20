package solution

import (
	"errors"
	"fmt"

	"github.com/mharbol/aoc-2023/util"
)

func GetSolution(day uint8) (Solution, error) {
	s, ok := allSolutions[day]
	if !ok {
		return nil, errors.New(fmt.Sprintf("No Solution exists for day %d", day))
	}
	return s, nil
}

type Solution interface {
	Day() uint8
	Part1([]string) (string, error)
	Part2([]string) (string, error)
}

func Solve(day, part uint8) (string, error) {
	s, err := GetSolution(day)
	if err != nil {
		return "", err
	}

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
		return "", errors.New(fmt.Sprintf("Part must be 1 or 2, received %d", part))
	}
}
