// this is just here as a template
package days

import (
	"fmt"

	"github.com/mharbol/aoc-2023/util"
)

type Day00 struct{}

func (s *Day00) Day() uint8 {
	return 0
}

func (s *Day00) Part1() (string, error) {
	lines, err := util.ReadProblemInfo(s.Day())
	if err != nil {
		return "", err
	}

	// solutions part 1 goes here
	fmt.Println(lines)

	return "", nil
}

func (s *Day00) Part2() (string, error) {
	lines, err := util.ReadProblemInfo(s.Day())
	if err != nil {
		return "", err
	}

	// solutions part 2 goes here
	fmt.Println(lines)

	return "", nil
}
