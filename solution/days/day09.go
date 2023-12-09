package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day09Part1(lines []string) (string, error) {

	acc := 0

	for _, line := range lines {
		acc += sumEnds(getDiffLevels(parseLineDay9(line)))
	}

	return fmt.Sprint(acc), nil
}

func Day09Part2(lines []string) (string, error) {

	acc := 0

	for _, line := range lines {
		acc += sumDiffFronts(getDiffLevels(parseLineDay9(line)))
	}

	return fmt.Sprint(acc), nil
}

func parseLineDay9(line string) []int {
	var nums []int
	for _, val := range strings.Split(line, " ") {
		number, _ := strconv.Atoi(val)
		nums = append(nums, number)
	}
	return nums
}

func getDiffs(nums []int) []int {
	var diffs []int
	for i := 1; i < len(nums); i++ {
		diffs = append(diffs, nums[i]-nums[i-1])
	}
	return diffs
}

func isAllZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func getDiffLevels(arr []int) [][]int {
	var levels [][]int

	for !isAllZeros(arr) {
		levels = append(levels, arr)
		arr = getDiffs(arr)
	}

	return levels
}

func sumEnds(arrs [][]int) int {
	tot := 0
	for _, arr := range arrs {
		tot += arr[len(arr)-1]
	}
	return tot
}

func sumDiffFronts(arrs [][]int) int {
	prev := 0

	for i := len(arrs) - 1; i >= 0; i-- {
        prev = arrs[i][0] - prev
	}
	return prev
}
