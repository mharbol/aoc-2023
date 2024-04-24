package day10

import (
	"fmt"
	"strconv"
)

func Part1(lines []string) string {

	s := findS(lines)

	vertexList = []coord{}

	prev := s
	curr := coord{row: prev.row + 1, col: prev.col}

	length := 1
	for lines[curr.row][curr.col] != 'S' {
		temp := curr
		curr = takeStep(curr, prev, lines)
		prev = temp
		length++
	}

	return fmt.Sprint(length/2 + length%2)
}

func Part2(lines []string) string {

	// traverse and mark pipes
	p1Ans := Part1(lines)
	halfLen, _ := strconv.Atoi(p1Ans)

	inside := shoelaceArea(vertexList) - halfLen

	return fmt.Sprint(inside)
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

var vertexList []coord

func takeStep(currPos, prevPos coord, lines []string) coord {

	switch lines[currPos.row][currPos.col] {
	case '|':
		if currPos.row > prevPos.row {
			return goSouth(currPos)
		} else {
			return goNorth(currPos)
		}
	case '-':
		if currPos.col > prevPos.col {
			return goEast(currPos)
		} else {
			return goWest(currPos)
		}
	case 'L':
		vertexList = append(vertexList, currPos)
		if currPos.row > prevPos.row {
			return goEast(currPos)
		} else {
			return goNorth(currPos)
		}
	case 'J':
		vertexList = append(vertexList, currPos)
		if currPos.row > prevPos.row {
			return goWest(currPos)
		} else {
			return goNorth(currPos)
		}
	case '7':
		vertexList = append(vertexList, currPos)
		if currPos.col > prevPos.col {
			return goSouth(currPos)
		} else {
			return goWest(currPos)
		}
	case 'F':
		vertexList = append(vertexList, currPos)
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
	return coord{row: currPos.row - 1, col: currPos.col}
}

func goSouth(currPos coord) coord {
	return coord{row: currPos.row + 1, col: currPos.col}
}

func goEast(currPos coord) coord {
	return coord{row: currPos.row, col: currPos.col + 1}
}

func goWest(currPos coord) coord {
	return coord{row: currPos.row, col: currPos.col - 1}
}

func shoelaceArea(vertices []coord) int {
	tot := 0
	length := len(vertices)
	for i := 0; i < length-1; i++ {
		tot += vertices[i].col*vertices[i+1].row - vertices[i].row*vertices[i+1].col
	}
	tot += vertices[length-1].col*vertices[0].row - vertices[length-1].row*vertices[0].col

	tot /= 2
	if tot >= 0 {
		return tot + 1
	} else {
		return -tot + 1
	}
}
