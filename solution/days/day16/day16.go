package day16

import (
	"fmt"
	"strings"
)

func Part1(lines []string) string {

	matrix := newMirrorMatrixFromStringSlice(lines)

	return fmt.Sprint(matrix.energyForInitialBeam(newBeam(0, 0, EAST)))
}

func Part2(lines []string) string {

	baseMatrix := newMirrorMatrixFromStringSlice(lines)

	numRows := len(lines)
	numCols := len(lines[0])

	maxEnergy := 0

	for startRow := 0; startRow < numRows; startRow++ {

		matrixEast := baseMatrix.shallowCopyForPart2()
		energyEast := matrixEast.energyForInitialBeam(newBeam(startRow, 0, EAST))

		matrixWest := baseMatrix.shallowCopyForPart2()
		energyWest := matrixWest.energyForInitialBeam(newBeam(startRow, numCols-1, WEST))

		maxNrg := max(energyEast, energyWest)
		if maxNrg > maxEnergy {
			maxEnergy = maxNrg
		}
	}

	for startCol := 0; startCol < numCols; startCol++ {

		matrixSouth := baseMatrix.shallowCopyForPart2()
		energySouth := matrixSouth.energyForInitialBeam(newBeam(0, startCol, SOUTH))

		matrixNorth := baseMatrix.shallowCopyForPart2()
		energyNorth := matrixNorth.energyForInitialBeam(newBeam(numRows-1, startCol, NORTH))

		maxNrg := max(energyNorth, energySouth)
		if maxNrg > maxEnergy {
			maxEnergy = maxNrg
		}
	}

	return fmt.Sprint(maxEnergy)
}

type direction int

// directions of travel
const (
	NORTH direction = 0
	SOUTH direction = 1
	EAST  direction = 2
	WEST  direction = 3
)

type mirrorMatrix struct {
	tiles          [][]byte
	isEnergized    [][]bool
	totalEnergy    int
	beamHistoryMap map[beam]bool
}

func newMirrorMatrixFromStringSlice(lines []string) *mirrorMatrix {
	var (
		tiles       [][]byte
		isEnergized [][]bool
	)

	for _, str := range lines {
		var row []byte
		var energizedRow []bool
		for _, char := range str {
			row = append(row, byte(char))
			energizedRow = append(energizedRow, false)
		}
		tiles = append(tiles, row)
		isEnergized = append(isEnergized, energizedRow)
	}

	return &mirrorMatrix{tiles: tiles, isEnergized: isEnergized, totalEnergy: 0, beamHistoryMap: make(map[beam]bool)}
}

func (mm *mirrorMatrix) energizeTile(b *beam) {
	if !mm.isEnergized[b.row][b.col] {
		mm.isEnergized[b.row][b.col] = true
		mm.totalEnergy++
	}
}

func (mm *mirrorMatrix) String() string {
	mirrors := ""
	energies := ""
	for _, row := range mm.tiles {
		for _, char := range row {
			mirrors += fmt.Sprint(string(char))
		}
		mirrors += "\n"
	}

	for _, row := range mm.isEnergized {
		for _, energy := range row {
			if energy {
				energies += "#"
			} else {
				energies += "."
			}
		}
		energies += "\n"
	}

	return strings.Trim(mirrors+"\n"+energies, "\n")
}

func (mm *mirrorMatrix) shallowCopyForPart2() *mirrorMatrix {

	var isEnergized [][]bool

	for row := 0; row < len(mm.isEnergized); row++ {
		var energizedRow []bool
		for col := 0; col < len(mm.isEnergized[row]); col++ {
			energizedRow = append(energizedRow, false)
		}
		isEnergized = append(isEnergized, energizedRow)
	}

	return &mirrorMatrix{tiles: mm.tiles, isEnergized: isEnergized, totalEnergy: 0, beamHistoryMap: make(map[beam]bool)}
}

// assesses the energization at a given tile and retuns a slice of beams to assess next
func (mm *mirrorMatrix) assessBeam(b *beam) []*beam {

	var beams []*beam

	// Exit case, out of bounds:
	if b.row < 0 || b.col < 0 || b.row >= len(mm.tiles) || b.col >= len(mm.tiles[b.row]) {
		return beams
	}

	_, ok := mm.beamHistoryMap[*b]
	if ok {
		return beams
	} else {
		mm.beamHistoryMap[*b] = true
	}

	mm.energizeTile(b)

	switch b.dir {
	case EAST:
		switch mm.tiles[b.row][b.col] {
		case '.':
			beams = append(beams, b.goEast())
			break
		case '-':
			beams = append(beams, b.goEast())
			break
		case '|':
			beams = append(beams, b.goNorth(), b.goSouth())
			break
		case '/':
			beams = append(beams, b.goNorth())
			break
		case '\\':
			beams = append(beams, b.goSouth())
			break
		default:
			panic("No case for EAST")
		}
		break
	case WEST:
		switch mm.tiles[b.row][b.col] {
		case '.':
			beams = append(beams, b.goWest())
			break
		case '-':
			beams = append(beams, b.goWest())
			break
		case '|':
			beams = append(beams, b.goNorth(), b.goSouth())
			break
		case '/':
			beams = append(beams, b.goSouth())
			break
		case '\\':
			beams = append(beams, b.goNorth())
			break
		default:
			panic("No case for WEST")
		}
		break
	case NORTH:
		switch mm.tiles[b.row][b.col] {
		case '.':
			beams = append(beams, b.goNorth())
			break
		case '|':
			beams = append(beams, b.goNorth())
			break
		case '-':
			beams = append(beams, b.goEast(), b.goWest())
			break
		case '/':
			beams = append(beams, b.goEast())
			break
		case '\\':
			beams = append(beams, b.goWest())
			break
		default:
			panic("No case for NORTH")
		}
		break
	case SOUTH:
		switch mm.tiles[b.row][b.col] {
		case '.':
			beams = append(beams, b.goSouth())
			break
		case '|':
			beams = append(beams, b.goSouth())
			break
		case '-':
			beams = append(beams, b.goEast(), b.goWest())
			break
		case '/':
			beams = append(beams, b.goWest())
			break
		case '\\':
			beams = append(beams, b.goEast())
			break
		default:
			panic("No case for SOUTH")
		}
		break
	default:
		panic(fmt.Sprintf("Case not defined for %v", *b))
	}

	return beams
}

func (mm *mirrorMatrix) energyForInitialBeam(initialBeam *beam) int {

	var beams []*beam = []*beam{initialBeam}
	var nextBeam *beam

	for len(beams) > 0 {
		beams, nextBeam = getPushedBeam(beams)
		beams = append(beams, mm.assessBeam(nextBeam)...)
	}
	return mm.totalEnergy
}

type beam struct {
	row int
	col int
	dir direction
}

func newBeam(row, col int, dir direction) *beam {
	return &beam{row: row, col: col, dir: dir}
}

func (b *beam) goNorth() *beam {
	return newBeam(b.row-1, b.col, NORTH)
}

func (b *beam) goSouth() *beam {
	return newBeam(b.row+1, b.col, SOUTH)
}

func (b *beam) goEast() *beam {
	return newBeam(b.row, b.col+1, EAST)
}

func (b *beam) goWest() *beam {
	return newBeam(b.row, b.col-1, WEST)
}

func getPushedBeam(beams []*beam) ([]*beam, *beam) {
	return beams[:len(beams)-1], beams[len(beams)-1]
}
