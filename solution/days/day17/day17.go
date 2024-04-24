package day17

import (
	"fmt"
	"strconv"
)

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

// purpose-built binary heap to see which of the traversed positions currently has the lowest heat loss
type positionPriorityQueue struct {
	heap []*lavaTraveralPosition
	size int `default:"0"`
}

func (q *positionPriorityQueue) push(position *lavaTraveralPosition) {
	q.heap = append(q.heap, position)

	idx := q.size

	for idx > 0 && q.heap[(idx-1)/2].heatLoss > q.heap[idx].heatLoss {
		q.swap((idx-1)/2, idx)
		idx = (idx - 1) / 2
	}

	q.size++
}

func (q *positionPriorityQueue) heapify(idx int) {
	if q.size <= 1 {
		return
	}

	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2
	smallestIdx := idx

	if leftIdx < q.size && q.heap[leftIdx].heatLoss < q.heap[smallestIdx].heatLoss {
		smallestIdx = leftIdx
	}
	if rightIdx < q.size && q.heap[rightIdx].heatLoss < q.heap[smallestIdx].heatLoss {
		smallestIdx = rightIdx
	}

	if smallestIdx != idx {
		q.swap(idx, smallestIdx)
		q.heapify(smallestIdx)
	}
}

func (q *positionPriorityQueue) pop() *lavaTraveralPosition {
	if q.size == 0 {
		return nil
	}
	lowest := q.heap[0]
	q.size--
	q.heap[0] = q.heap[q.size]
	q.heap = q.heap[:q.size]
	q.heapify(0)

	return lowest
}

func (q *positionPriorityQueue) swap(i, j int) {
	temp := q.heap[i]
	q.heap[i] = q.heap[j]
	q.heap[j] = temp
}

type lavaTraveralMatrix struct {
	// heat loss by row/col position
	heatLossMatrix [][]int
	// positions to assess
	positions *positionPriorityQueue
	// visited positions
	visited map[lavaCoordinate]int
	// sizes for rows and columns
	rowSize int
	colSize int
}

func newLavaTraversalMatrixFromStrings(lines []string) *lavaTraveralMatrix {
	var heatLoss [][]int

	for _, line := range lines {
		var rowHeatLoss []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			rowHeatLoss = append(rowHeatLoss, num)
		}
		heatLoss = append(heatLoss, rowHeatLoss)
	}

	return &lavaTraveralMatrix{heatLossMatrix: heatLoss, positions: &positionPriorityQueue{},
		visited: make(map[lavaCoordinate]int), rowSize: len(heatLoss), colSize: len(heatLoss[0])}
}

func (ltm *lavaTraveralMatrix) String() string {
	out := ""
	for _, row := range ltm.heatLossMatrix {
		out += fmt.Sprintln(row)
	}
	return out
}

func (ltm *lavaTraveralMatrix) walkPosition(position *lavaTraveralPosition, dir walkDirection) {
	if position.lastDirection == -dir || (position.lastDirection == dir && position.stepsTakenInDir >= MAX_STEPS_NORMAL) {
		return
	}

	posCopy := position.copy()

	switch dir {
	case UP:
		if position.row == 0 {
			return
		}
		posCopy.row--
		break
	case DOWN:
		if position.row >= ltm.rowSize-1 {
			return
		}
		posCopy.row++
		break
	case RIGHT:
		if position.col >= ltm.colSize-1 {
			return
		}
		posCopy.col++
		break
	case LEFT:
		if position.col == 0 {
			return
		}
		posCopy.col--
		break
	}

	if posCopy.lastDirection == dir {
		posCopy.stepsTakenInDir++
	} else {
		posCopy.lastDirection = dir
		posCopy.stepsTakenInDir = 1
	}

	// Check if been here before
	currPos := lavaCoordinate{row: posCopy.row, col: posCopy.col, dir: posCopy.lastDirection, steps: posCopy.stepsTakenInDir}
	_, ok := ltm.visited[currPos]
	if ok {
		return
	} else {
		ltm.visited[currPos] = posCopy.heatLoss
	}

	posCopy.heatLoss += ltm.heatLossMatrix[posCopy.row][posCopy.col]

	ltm.positions.push(posCopy)
}

// for this case, we assume we have already exhaused all other walking cases in this direction and they are already in the queue
// so we will only make a r/l turn at this point and go through all walking scenarios from there
func (ltm *lavaTraveralMatrix) walkUltraCrucible(position *lavaTraveralPosition, dir walkDirection) {

	// If I have dealt with this position before in a more efficient way...
	currPos := lavaCoordinate{row: position.row, col: position.col, dir: position.lastDirection, steps: 0}
	heatLoss, ok := ltm.visited[currPos]
	if ok {
		if heatLoss < position.heatLoss {
			return
		} else {
			ltm.visited[currPos] = position.heatLoss
		}
	} else {
		ltm.visited[currPos] = position.heatLoss
	}

	// verify valid direction
	if position.lastDirection == dir || position.lastDirection == -dir {
		return
	}

	// Shows direction to walk in the next steps
	rowWalkVal := 0
	colWalkVal := 0

	// Verify direction is within min walking distance
	switch dir {
	case UP:
		if position.row < MIN_STEPS_ULTRA {
			return
		} else {
			rowWalkVal = -1
		}
		break
	case DOWN:
		if position.row > ltm.rowSize-MIN_STEPS_ULTRA {
			return
		} else {
			rowWalkVal = 1
		}
		break
	case RIGHT:
		if position.col > ltm.colSize-MIN_STEPS_ULTRA {
			return
		} else {
			colWalkVal = 1
		}
		break
	case LEFT:
		if position.col < MIN_STEPS_ULTRA {
			return
		} else {
			colWalkVal = -1
		}
		break
	}

	// walk one less than the initial min steps in the correct direction
	posCopy := position.copy()
	posCopy.lastDirection = dir
	posCopy.stepsTakenInDir = 0
	for i := 0; i < MIN_STEPS_ULTRA-1; i++ {
		posCopy.row += rowWalkVal
		posCopy.col += colWalkVal
		posCopy.heatLoss += ltm.heatLossMatrix[posCopy.row][posCopy.col]
		posCopy.stepsTakenInDir++
	}

	// continue making steps in the given direction until cannot or hit a cached position
	currCopy := posCopy
	for i := 1; i <= MAX_STEPS_ULTRA-MIN_STEPS_ULTRA+1; i++ {
		currCopy = currCopy.copy()

		currCopy.row += rowWalkVal
		currCopy.col += colWalkVal
		currCopy.stepsTakenInDir++

		if currCopy.stepsTakenInDir > MAX_STEPS_ULTRA || currCopy.row < 0 || currCopy.row >= ltm.rowSize || currCopy.col < 0 || currCopy.col >= ltm.colSize {
			return
		}

		currCopy.heatLoss += ltm.heatLossMatrix[currCopy.row][currCopy.col]

		ltm.positions.push(currCopy)
	}
}
