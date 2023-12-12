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

	return "", nil
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
