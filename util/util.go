package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/mharbol/aoc-2023/solution"
)

const _INPUT_PATH string = "../inputs/"

var dayInputMap map[uint8][]string = make(map[uint8][]string)

func ReadProblemInfo(s solution.Solution) ([]string, error) {

	var fileContents []string
	var ok bool
	var err error

	fileContents, ok = dayInputMap[s.Day()]

	if !ok {
		fileContents, err = readDayInput(makeFilePath(s.Day()))
		if err != nil {
			return nil, err
		}
		dayInputMap[s.Day()] = fileContents
	}
	return fileContents, nil
}

func readDayInput(filePath string) ([]string, error) {
	fileDump, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileDump), "\n"), nil
}

func makeFilePath(day uint8) string {
	return fmt.Sprintf("%sday_%02d.txt", _INPUT_PATH, day)
}
