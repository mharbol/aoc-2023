package days

import "fmt"

func Day11Part1(lines []string) (string, error) {

	acc := 0

	universe := expandUniverse(lines)
	galxies := getGalaxies(universe)

	for x := 0; x < len(galxies)-1; x++ {
		for y := x + 1; y < len(galxies); y++ {
			acc += manhattanDistance(galxies[x], galxies[y])
		}
	}

	return fmt.Sprint(acc), nil
}

func Day11Part2(lines []string) (string, error) {

	acc := 0

	const EXPANSION = 1000000

	universe, blankRows, blankCols := observeUniverse(lines)
	galaxies := getGalaxies(universe)

	for x := 0; x < len(galaxies)-1; x++ {
		for y := x + 1; y < len(galaxies); y++ {
			acc += expandedDistance(galaxies[x], galaxies[y], blankRows, blankCols, EXPANSION)
		}
	}

	return fmt.Sprint(acc), nil
}

type galaxyCoord struct {
	row int
	col int
}

func i32Abs(n int32) int32 {
	var y int32 = (n >> 31)
	return (n ^ y) - y
}

func expandUniverse(lines []string) [][]byte {

	// get expanded columns
	doubleColSet := make(map[int]bool)
	for col := 0; col < len(lines[0]); col++ {
		allDots := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] == '#' {
				allDots = false
				break
			}
		}
		if allDots {
			doubleColSet[col] = true
		}
	}

	// loop over and make rows
	var out [][]byte
	for row := 0; row < len(lines); row++ {
		var line []byte
		allDots := true
		for col := 0; col < len(lines[row]); col++ {
			_, ok := doubleColSet[col]
			if ok {
				line = append(line, '.')
			}
			if lines[row][col] != '.' {
				allDots = false
			}
			line = append(line, lines[row][col])
		}
		out = append(out, line)
		if allDots {
			out = append(out, line)
		}
	}

	return out
}

func getGalaxies(universe [][]byte) []galaxyCoord {
	var out []galaxyCoord
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col] == '#' {
				out = append(out, galaxyCoord{row: row, col: col})
			}
		}
	}
	return out
}

func manhattanDistance(g1, g2 galaxyCoord) int {
	return int(i32Abs(int32(g1.col)-int32(g2.col))) + int(i32Abs(int32(g1.row)-int32(g2.row)))
}

func observeUniverse(lines []string) ([][]byte, map[int]bool, map[int]bool) {

	blankCols := make(map[int]bool)
	blankRows := make(map[int]bool)
	var universe [][]byte

	for col := 0; col < len(lines[0]); col++ {
		allDots := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] == '#' {
				allDots = false
				break
			}
		}
		if allDots {
			blankCols[col] = true
		}
	}

	for row := 0; row < len(lines); row++ {
		allDots := true
		var line []byte
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] != '.' {
				allDots = false
			}
			line = append(line, lines[row][col])
		}
		universe = append(universe, line)
		if allDots {
			blankRows[row] = true
		}
	}

	return universe, blankRows, blankCols
}

func expandedDistance(x, y galaxyCoord, blankRows, blankCols map[int]bool, exp int) int {

	var rowTop, rowBot, colLeft, colRight int

	if x.row > y.row {
		rowTop = y.row
		rowBot = x.row
	} else {
		rowTop = x.row
		rowBot = y.row
	}
	if x.col > y.col {
		colRight = x.col
		colLeft = y.col
	} else {
		colRight = y.col
		colLeft = x.col
	}

	expCount := 0
	for i := rowTop; i < rowBot; i++ {
		_, ok := blankRows[i]
		if ok {
			expCount++
		}
	}
	for i := colLeft; i < colRight; i++ {
		_, ok := blankCols[i]
		if ok {
			expCount++
		}
	}
	return rowBot - rowTop + colRight - colLeft + (exp-1)*expCount
}
