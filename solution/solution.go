package solution

import (
	"errors"
	"fmt"

	"github.com/mharbol/aoc-2023/util"
)

type solFunc func([]string) string

type solFuncTuple struct {
	part1 solFunc
	part2 solFunc
}

func Solve(day, part uint8) (string, error) {

	sol, err := GetSolution(day)
	if err != nil {
		return "", err
	}

	lines, err := util.GetDayInput(day)
	if err != nil {
		return "", err
	}

	switch part {
	case 1:
		return sol.part1(lines), nil
	case 2:
		return sol.part2(lines), nil
	default:
		return "", errors.New(fmt.Sprintf("Part must be 1 or 2, received %d", part))
	}
}

func GetSolution(day uint8) (*solFuncTuple, error) {
	sol, ok := allSolutions[day]
	if !ok {
		return nil, errors.New(fmt.Sprintf("No Solution exists for day %d", day))
	}
	return sol, nil
}
