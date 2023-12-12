package days

import "fmt"

func Day10Part1(lines []string) (string, error) {

	s := findS(lines)
	fmt.Println(s.row, s.col)

    prev := s
    curr := coord{row: prev.row + 1, col: prev.col}

    length := 1
    for lines[curr.row][curr.col] != 'S' {
        fmt.Println("curr: ", curr, " prev: ", prev)
        temp := curr
        curr = takeStep(curr, prev, lines)
        prev = temp
        length++
    }

	return fmt.Sprint(length / 2 + length % 2), nil
}

func Day10Part2(lines []string) (string, error) {

	return "", nil
}

type coord struct {
	row int
	col int
}

func findS(lines []string) coord {
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] == 'S' {
				return coord{row: row, col: col}
			}
		}
	}
	return coord{-1, -1}
}

func takeStep(currPos, prevPos coord, lines []string) coord {

	switch lines[currPos.row][currPos.col] {
	case '|':
		if currPos.row > prevPos.row { // go south
			return goSouth(currPos)
		} else { // go north
			return goNorth(currPos)
		}
	case '-':
		if currPos.col > prevPos.col { // go east
			return goEast(currPos)
		} else { // go west
			return goWest(currPos)
		}
	case 'L':
		if currPos.row > prevPos.row { // go east
			return goEast(currPos)
		} else { // go north
			return goNorth(currPos)
		}
	case 'J':
		if currPos.row > prevPos.row { // go west
			return goWest(currPos)
		} else {
			return goNorth(currPos)
		}
    case '7':
        if currPos.col > prevPos.col {
            return goSouth(currPos)
        } else {
            return goWest(currPos)
        }
    case 'F':
        if currPos.row < prevPos.row {
            return goEast(currPos)
        } else {
            return goSouth(currPos)
        }
    default:
        fmt.Println("Came to a ", lines[currPos.row][currPos.col], " at ", currPos)
	}

	return coord{-1, -1}
}

func goNorth(currPos coord) coord {
    fmt.Println("go north")
	return coord{row: currPos.row - 1, col: currPos.col}
}

func goSouth(currPos coord) coord {
    fmt.Println("go south")
	return coord{row: currPos.row + 1, col: currPos.col}
}

func goEast(currPos coord) coord {
    fmt.Println("go east")
	return coord{row: currPos.row, col: currPos.col + 1}
}

func goWest(currPos coord) coord {
    fmt.Println("go west")
	return coord{row: currPos.row, col: currPos.col - 1}
}
