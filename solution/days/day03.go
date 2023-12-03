package days

import "fmt"

func Day03Part1(lines []string) (string, error) {

	acc := 0

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {

			// see if we are on a number
			if isDigit(lines[row][col]) {
				endOfNum := endOfNumber(col, lines[row])

				// check box and parse number as needed
				if checkBoxForSymbols(row, col-1, endOfNum, lines) {
					acc += parseNumber(col, lines[row])
				}

				// progress col to the end of the number
				col = endOfNum
			}
		}
	}

	return fmt.Sprint(acc), nil
}

func Day03Part2(lines []string) (string, error) {

	acc := 0

	// keeps track of locations of stars and their neighboring numbers
	starsWithNums := make(map[intPair][]int)

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {

			// check to see if the number has a neighboring '*'
			if isDigit(lines[row][col]) {

				endOfNum := endOfNumber(col, lines[row])
				num := parseNumber(col, lines[row])
				stars := checkBoxForStars(row, col-1, endOfNum, lines)

				if stars != nil {
					for _, location := range stars {
						valsAdjacentToStar := starsWithNums[location]
						valsAdjacentToStar = append(valsAdjacentToStar, num)
						starsWithNums[location] = valsAdjacentToStar
					}
				}

				col = endOfNum
			}
		}
	}
	for _, value := range starsWithNums {
		if len(value) == 2 {
			acc += value[0] * value[1]
		}
	}

	return fmt.Sprint(acc), nil
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isSymbol(char byte) bool {
	return !isDigit(char) && char != '.'
}

// index of the character after the current number
func endOfNumber(col int, line string) int {
	for col < len(line) && isDigit(line[col]) {
		col++
	}
	return col
}

func checkBoxForSymbols(baseRow, incStart, incEnd int, lines []string) bool {
	// check left
	if incStart >= 0 {
		if isSymbol(lines[baseRow][incStart]) {
			return true
		}
	} else {
		incStart = 0
	}
	// check right
	if incEnd < len(lines[baseRow]) {
		if isSymbol(lines[baseRow][incEnd]) {
			return true
		}
	} else {
		incEnd = len(lines[baseRow]) - 1
	}
	// check above
	if baseRow > 0 {
		for i := incStart; i <= incEnd; i++ {
			if isSymbol(lines[baseRow-1][i]) {
				return true
			}
		}
	}
	// check below
	if baseRow+1 < len(lines) {
		for i := incStart; i <= incEnd; i++ {
			if isSymbol(lines[baseRow+1][i]) {
				return true
			}
		}
	}
	return false
}

func parseNumber(col int, line string) int {
	out := 0

	for col < len(line) && isDigit(line[col]) {
		out = out*10 + int(line[col]-'0')
		col++
	}
	return out
}

type intPair struct {
	row int
	col int
}

func isStar(char byte) bool {
	return '*' == char
}

func checkBoxForStars(baseRow, incStart, incEnd int, lines []string) []intPair {

	var out []intPair

	// check left
	if incStart >= 0 {
		if isStar(lines[baseRow][incStart]) {
			out = append(out, intPair{baseRow, incStart})
		}
	} else {
		incStart = 0
	}
	// check right
	if incEnd < len(lines[baseRow]) {
		if isStar(lines[baseRow][incEnd]) {
			out = append(out, intPair{baseRow, incEnd})
		}
	} else {
		incEnd = len(lines[baseRow]) - 1
	}
	// check above
	if baseRow > 0 {
		for i := incStart; i <= incEnd; i++ {
			if isStar(lines[baseRow-1][i]) {
				out = append(out, intPair{baseRow - 1, i})
			}
		}
	}
	// check below
	if baseRow+1 < len(lines) {
		for i := incStart; i <= incEnd; i++ {
			if isStar(lines[baseRow+1][i]) {
				out = append(out, intPair{baseRow + 1, i})
			}
		}
	}
	return out
}
