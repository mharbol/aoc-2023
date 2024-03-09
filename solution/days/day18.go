package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day18Part1(lines []string) (string, error) {

	digDirs := parseDigPatternP1(lines)

	return fmt.Sprint(getFullArea(digDirs)), nil
}

func Day18Part2(lines []string) (string, error) {

	digDirs := parseDigPatternP2(lines)

	return fmt.Sprint(getFullArea(digDirs)), nil
}

type digDirection struct {
	dir    byte
	length int64
}

type trenchCoord struct {
	row int64
	col int64
}

func parseDigPatternP1(lines []string) []*digDirection {
	var digDirections []*digDirection
	for _, line := range lines {
		elements := strings.Split(line, " ")

		length, _ := strconv.Atoi(elements[1])

		digDirections = append(digDirections, &digDirection{dir: elements[0][0], length: int64(length)})
	}
	return digDirections
}

func parseDigPatternP2(lines []string) []*digDirection {
	var digDirections []*digDirection
	for _, line := range lines {
		elements := strings.Split(line, " ")
		length, _ := strconv.ParseInt(elements[2][2:7], 16, 32)

		var dirByte byte
		switch elements[2][7] {
		case '0':
			dirByte = 'R'
			break
		case '1':
			dirByte = 'D'
			break
		case '2':
			dirByte = 'L'
			break
		case '3':
			dirByte = 'U'
			break
		}

		digDirections = append(digDirections, &digDirection{dir: dirByte, length: length})
	}
	return digDirections
}

func shoelaceTrenchArea(points []*trenchCoord) int64 {
	var tot int64 = 0
	length := len(points)

	for i := 0; i < length-1; i++ {
		tot += points[i].col*points[i+1].row - points[i].row*points[i+1].col
	}

	tot /= 2

	if tot >= 0 {
		return tot
	} else {
		return -tot
	}
}

func getFullArea(digDirs []*digDirection) int64 {

	var coords []*trenchCoord
	currCoord := &trenchCoord{0, 0}
	coords = append(coords, currCoord)
	var perimeter int64 = 0
	for _, digDir := range digDirs {
		currCoord = &trenchCoord{row: currCoord.row, col: currCoord.col}
		switch digDir.dir {
		case 'U':
			currCoord.row += digDir.length
			break
		case 'D':
			currCoord.row -= digDir.length
			break
		case 'R':
			currCoord.col += digDir.length
			break
		case 'L':
			currCoord.col -= digDir.length
			break
		}
		perimeter += digDir.length
		coords = append(coords, currCoord)
	}

	area := shoelaceTrenchArea(coords)

	return area + perimeter/2 + 1
}
