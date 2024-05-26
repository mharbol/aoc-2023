package day21

type coord struct {
	row, col int
}

func walkToPoint(point coord, walkMap map[coord]int, steps int) {
	stepsCurrent, ok := walkMap[point]
	if !ok || steps%2 == 0 && stepsCurrent%2 == 1 {
		walkMap[point] = steps
	}
}

func walkFromPoint(point coord, lines []string, walkMap map[coord]int, stepNumber int) {
    // North
    if point.row-1 >= 0 && lines[point.row-1][point.col] != '#'{
        walkToPoint(coord{row: point.row-1, col: point.col}, walkMap, stepNumber)
    }
    // West
    if point.col-1 >= 0 && lines[point.row][point.col-1] != '#'{
        walkToPoint(coord{row: point.row, col: point.col-1}, walkMap, stepNumber)
    }
    // South
    if point.row+1 < len(lines) && lines[point.row+1][point.col] != '#' {
        walkToPoint(coord{row: point.row+1, col: point.col}, walkMap, stepNumber)
    }
    // East
    if point.col+1 < len(lines[point.row]) && lines[point.row][point.col+1] != '#' {
        walkToPoint(coord{row: point.row, col: point.col+1}, walkMap, stepNumber)
    }
}

func findStart(lines []string) coord {
    for y, row := range lines {
        for x, col := range row {
            if col == 'S' {
                return coord{row: y, col: x}
            }
        }
    }
    return coord{row: -1, col: -1}
}
