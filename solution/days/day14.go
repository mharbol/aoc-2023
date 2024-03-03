package days

import "fmt"

func Day14Part1(lines []string) (string, error) {

	plat := linesToPlatform(lines)
	plat.tiltNorth()

	return fmt.Sprint(plat.sumLoad()), nil
}

func Day14Part2(lines []string) (string, error) {

	plat := linesToPlatform(lines)

	// uses the values as keys and stores which step the platform reached that state
	valsAtStep := make(map[string]int)

	step := 0
    var cycleStart int
    
	valsAtStep[plat.String()] = step
	for {
		step++
		plat.cycle()
		key := plat.String()
		num, ok := valsAtStep[key]
		if !ok {
			valsAtStep[key] = step
		} else {
            // once here, we have hit a cycle in the platform states
            cycleStart = num
			break
		}
	}

    cyclesToCompleteOnceInCycle := 1_000_000_000 - cycleStart
    cyclicPeriod := step - cycleStart
    remainingSteps := cyclesToCompleteOnceInCycle % cyclicPeriod

    for i := 0; i < remainingSteps; i++ {
        plat.cycle()
    }

	return fmt.Sprint(plat.sumLoad()), nil
}

const (
	EMPTY = 0
	CUBE  = 1
	ROUND = 2
)

type platform struct {
	rows    [][]int
	numRows int
	numCols int
}

func (p *platform) cycle() {
	p.tiltNorth()
	p.tiltWest()
	p.tiltSouth()
	p.tiltEast()
}

func (p *platform) tiltNorth() {
	for idx := 1; idx < len(p.rows); idx++ {
		p.moveRowNorth(idx)
	}
}

func (p *platform) tiltSouth() {
	for idx := p.numRows - 1; idx >= 0; idx-- {
		p.moveRowSouth(idx)
	}
}

func (p *platform) tiltWest() {

	for _, row := range p.rows {
		for col := 0; col < p.numCols-1; col++ {
			if row[col] != EMPTY {
				continue
			}

			// find the next ROUND or CUBE
			rightPtr := col + 1

			for rightPtr+1 < p.numCols && row[rightPtr] == EMPTY {
				rightPtr++
			}
			if row[rightPtr] == ROUND {
				row[col] = ROUND
				row[rightPtr] = EMPTY
			} else if row[rightPtr] == CUBE {
				col = rightPtr
			}
		}
	}
}

func (p *platform) tiltEast() {

	for _, row := range p.rows {
		for col := p.numCols - 1; col > 0; col-- {
			if row[col] != EMPTY {
				continue
			}

			// find the next ROUND or CUBE
			leftPtr := col - 1

			for leftPtr-1 >= 0 && row[leftPtr] == EMPTY {
				leftPtr--
			}
			if row[leftPtr] == ROUND {
				row[col] = ROUND
				row[leftPtr] = EMPTY
			} else if row[leftPtr] == CUBE {
				col = leftPtr
			}
		}
	}
}

func (p *platform) moveRowNorth(rowIdx int) {
	for col, space := range p.rows[rowIdx] {
		if space == ROUND {
			// find northmost spot
			northIdx := rowIdx - 1
			for northIdx >= 0 && p.rows[northIdx][col] == EMPTY {
				northIdx--
			}
			if northIdx+1 != rowIdx {
				p.rows[northIdx+1][col] = ROUND
				p.rows[rowIdx][col] = EMPTY
			}
		}
	}
}

func (p *platform) moveRowSouth(rowIdx int) {
	for col, space := range p.rows[rowIdx] {
		if space == ROUND {
			// find southmost spot
			southIdx := rowIdx + 1
			for southIdx < p.numRows && p.rows[southIdx][col] == EMPTY {
				southIdx++
			}
			if southIdx-1 != rowIdx {
				p.rows[southIdx-1][col] = ROUND
				p.rows[rowIdx][col] = EMPTY
			}
		}
	}
}

func (p *platform) countRoundsForRow(row int) int {
	ret := 0
	for _, rock := range p.rows[row] {
		if rock == ROUND {
			ret++
		}
	}
	return ret
}

func (p *platform) sumLoad() int {

	tot := 0
	length := len(p.rows)
	for idx := 0; idx < length; idx++ {
		tot += p.countRoundsForRow(idx) * (length - idx)
	}

	return tot
}

func (p *platform) String() string {
	out := ""

	for _, row := range p.rows {
		for _, col := range row {
			out += fmt.Sprint(col)
		}
	}

	return out
}

func linesToPlatform(lines []string) *platform {

	var rows [][]int

	for _, line := range lines {
		var row []int
		for _, char := range line {
			switch char {
			case 'O':
				row = append(row, ROUND)
				break
			case '#':
				row = append(row, CUBE)
				break
			case '.':
				row = append(row, EMPTY)
				break
			}
		}
		rows = append(rows, row)
	}

	return &platform{rows: rows, numRows: len(rows), numCols: len(rows[0])}
}
