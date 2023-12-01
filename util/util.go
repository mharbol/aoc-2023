package util

import (
	"fmt"
	"os"
	"strings"
)

const _INPUT_PATH string = "inputs/"

var dayInputMap map[uint8][]string = make(map[uint8][]string)

func GetDayInput(day uint8) ([]string, error) {

	var (
		fileContents []string
		ok           bool
		err          error
	)

	fileContents, ok = dayInputMap[day]

	if !ok {
		fileContents, err = readDayInput(makeFilePath(day))
		if err != nil {
			return nil, err
		}
		dayInputMap[day] = fileContents
	}
	return fileContents, nil
}

func readDayInput(filePath string) ([]string, error) {
	fileDump, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
    // when I split over \n, the last item is a ""
    // prune it off
    dumpSlice := strings.Split(string(fileDump), "\n")
    if len(dumpSlice) > 1 {
        dumpSlice = dumpSlice[:len(dumpSlice) - 1]
    }
	return dumpSlice, nil
}

func makeFilePath(day uint8) string {
	return fmt.Sprintf("%sday_%02d.txt", _INPUT_PATH, day)
}
