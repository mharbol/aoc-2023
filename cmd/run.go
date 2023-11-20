package cmd

import "github.com/mharbol/aoc-2023/solution"

func RunProb(parser ArgParser) (string, error) {

	parser.Parse()

	day, part := parser.GetDay(), parser.GetPart()

	out, err := solution.Solve(uint8(day), uint8(part))

	if err != nil {
		return "", err
	}

	return out, nil
}
