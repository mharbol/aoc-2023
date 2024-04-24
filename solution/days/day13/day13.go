package day13

import "fmt"

func Part1(lines []string) string {

	totVert := 0
	totHorz := 0

	for _, block := range getBlocks(lines) {
		totVert += getReflectionIndex(getVertIntRepr(block))
		totHorz += getReflectionIndex(getHorzIntRepr(block))
	}

	return fmt.Sprint(totHorz*100 + totVert)
}

func Part2(lines []string) string {

	totVert := 0
	totHorz := 0

	for _, block := range getBlocks(lines) {
		totVert += getSmudgeReflectionIndex(getVertIntRepr(block))
		totHorz += getSmudgeReflectionIndex(getHorzIntRepr(block))
	}

	return fmt.Sprint(totHorz*100 + totVert)
}

func getBlocks(lines []string) [][]string {

	var strings [][]string
	var currSlice []string

	for _, line := range lines {
		if line == "" {
			strings = append(strings, currSlice)
			currSlice = []string{}
		} else {
			currSlice = append(currSlice, line)
		}
	}

	strings = append(strings, currSlice)

	return strings
}

func getHorzIntRepr(block []string) []uint32 {

	var rows []uint32

	for _, row := range block {
		var num uint32 = 0
		for _, char := range row {
			if char == '#' {
				num |= 1
			}
			num <<= 1
		}
		rows = append(rows, num)
	}
	return rows
}

func getVertIntRepr(block []string) []uint32 {

	var cols []uint32

	for col := 0; col < len(block[0]); col++ {
		var num uint32 = 0
		for row := 0; row < len(block); row++ {
			if block[row][col] == '#' {
				num |= 1
			}
			num <<= 1
		}
		cols = append(cols, num)
	}

	return cols
}

// retuns the index of the reflection, 0 if none
func getReflectionIndex(arr []uint32) int {
	for idx := 0; idx < len(arr)-1; idx++ {
		if isReflectionIndex(idx, arr) {
			return idx + 1
		}
	}
	return 0
}

// retuns the index of the smudge reflection, 0 if none
func getSmudgeReflectionIndex(arr []uint32) int {
	for idx := 0; idx < len(arr)-1; idx++ {
		if isSmudgeReflectionIndex(idx, arr) {
			return idx + 1
		}
	}
	return 0
}

func isReflectionIndex(idx int, arr []uint32) bool {

	left := idx
	right := idx + 1

	for left >= 0 && right < len(arr) {
		if arr[left] != arr[right] {
			return false
		}
		left--
		right++
	}

	return true
}

func isSmudgeReflectionIndex(idx int, arr []uint32) bool {

	left := idx
	right := idx + 1

	powerOfTwoCount := 0

	for left >= 0 && right < len(arr) {
		xor := arr[left] ^ arr[right]
		pow2 := isPowerOf2(xor)
		if arr[left] != arr[right] && !pow2 {
			return false
		}
		if pow2 {
			powerOfTwoCount += 1
		}
		left--
		right++
	}

	return powerOfTwoCount == 1

}

func isPowerOf2(num uint32) bool {
	return num != 0 && num&(num-1) == 0
}
