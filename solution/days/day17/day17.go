package day17

import "fmt"

func Part1(lines []string) string {

	ltm := newLavaTraversalMatrixFromStrings(lines)

	nextPos := &lavaTraveralPosition{row: 0, col: 0, heatLoss: 0, lastDirection: 0, stepsTakenInDir: 0}
	ltm.positions.push(nextPos)

	endRow := len(ltm.heatLossMatrix) - 1
	endCol := len(ltm.heatLossMatrix[endRow]) - 1

	for nextPos.row != endRow || nextPos.col != endCol {
		ltm.walkPosition(nextPos, UP)
		ltm.walkPosition(nextPos, DOWN)
		ltm.walkPosition(nextPos, RIGHT)
		ltm.walkPosition(nextPos, LEFT)
		nextPos = ltm.positions.pop()
	}

	return fmt.Sprint(nextPos.heatLoss)
}

func Part2(lines []string) string {

	ltm := newLavaTraversalMatrixFromStrings(lines)
	nextPos := &lavaTraveralPosition{row: 0, col: 0, heatLoss: 0, lastDirection: 0, stepsTakenInDir: 0}

	endRow := len(ltm.heatLossMatrix) - 1
	endCol := len(ltm.heatLossMatrix[endRow]) - 1

	for nextPos.row != endRow || nextPos.col != endCol {
		ltm.walkUltraCrucible(nextPos, UP)
		ltm.walkUltraCrucible(nextPos, DOWN)
		ltm.walkUltraCrucible(nextPos, RIGHT)
		ltm.walkUltraCrucible(nextPos, LEFT)
		nextPos = ltm.positions.pop()
	}

	return fmt.Sprint(nextPos.heatLoss)
}

const (
	MAX_STEPS_NORMAL = 3
	MIN_STEPS_ULTRA  = 4
	MAX_STEPS_ULTRA  = 10
)

type walkDirection int

const (
	UP    walkDirection = -1
	DOWN  walkDirection = 1
	RIGHT walkDirection = -2
	LEFT  walkDirection = 2
)

type lavaTraveralPosition struct {
	row             int
	col             int
	heatLoss        int
	lastDirection   walkDirection
	stepsTakenInDir int
}

func (ltp *lavaTraveralPosition) String() string {
	return fmt.Sprintf("{heatLoss: %d, row: %d, col: %d, lastDir: %d, steps: %d}", ltp.heatLoss, ltp.row, ltp.col, ltp.lastDirection, ltp.stepsTakenInDir)
}

func (ltp *lavaTraveralPosition) copy() *lavaTraveralPosition {
	return &lavaTraveralPosition{row: ltp.row, col: ltp.col, heatLoss: ltp.heatLoss, lastDirection: ltp.lastDirection, stepsTakenInDir: ltp.stepsTakenInDir}
}

type lavaCoordinate struct {
	row   int
	col   int
	dir   walkDirection
	steps int
}
