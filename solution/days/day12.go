package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day12Part1(lines []string) (string, error) {

	acc := 0

	// clear the cache
	cache = make(map[cacheKey]int)

	for _, line := range lines {
		acc += countArrangements(parseDay12(line))
	}

	return fmt.Sprint(acc), nil
}

func Day12Part2(lines []string) (string, error) {

	acc := 0

	// clear the cache
	cache = make(map[cacheKey]int)

	for _, line := range lines {
		acc += countArrangements(unfold(parseDay12(line)))
	}

	return fmt.Sprint(acc), nil
}

type cacheKey struct {
	line   string
	counts string // cannot use slice as part of key
}

func putCache(line string, counts []int, count int) {
	val := ""
	for _, count := range counts {
		val += fmt.Sprint(count) + " "
	}
	key := cacheKey{line: line, counts: val}
	cache[key] = count
}

func getCache(line string, counts []int) (int, bool) {
	val := ""
	for _, count := range counts {
		val += fmt.Sprint(count) + " "
	}
	key := cacheKey{line: line, counts: val}
	num, ok := cache[key]
	return num, ok
}

var cache map[cacheKey]int

func parseDay12(line string) (string, []int) {
	info := strings.Split(line, " ")
	nums := strings.Split(info[1], ",")
	var numsSlice []int
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		numsSlice = append(numsSlice, num)
	}
	return info[0], numsSlice
}

func countArrangements(line string, counts []int) int {

	// check cache
	if num, ok := getCache(line, counts); ok {
		return num
	}

	// base case, out of counts
	if len(counts) == 0 {
		// no more hashes
		if strings.Index(line, "#") == -1 {
			putCache(line, counts, 1)
			return 1
		} else {
			// at least one more hash
			putCache(line, counts, 0)
			return 0
		}
	}

	// base case, no more characters but still counts
	if len(line) == 0 {
		putCache(line, counts, 0)
		return 0
	}

	// building to recursive step
	// look at the next element for counts and line
	nextChar := line[0]
	nextCount := counts[0]

	// logic to be used if the next char is a dot
	processDot := func() int {
		// skip over and look for next hash
		return countArrangements(line[1:], counts)
	}

	processHash := func() int {
		// if the first item is a hash, then the next n chars have
		// to be treated as hashes
		if len(line) < nextCount {
			putCache(line, counts, 0)
			return 0
		}
		chunk := line[:nextCount]
		chunk = strings.ReplaceAll(chunk, "?", "#")
		// if not all hashes or too short, then leave
		if strings.Index(chunk, ".") != -1 {
			putCache(line, counts, 0)
			return 0
		}

		// if the rest of the line is the last count, then left with valid or not
		if len(line) == nextCount {
			// valid
			if len(counts) == 1 {
				putCache(line, counts, 1)
				return 1
			} else {
				putCache(line, counts, 0)
				return 0
			}
		}

		// check that the next char is a separator and not a hash
		if line[nextCount] == '.' || line[nextCount] == '?' {
			return countArrangements(line[nextCount+1:], counts[1:])
		}
		// no possibilities
		return 0
	}

	var out int
	if nextChar == '#' {
		out = processHash()
	} else if nextChar == '.' {
		out = processDot()
	} else {
		out = processDot() + processHash()
	}
	putCache(line, counts, out)
	return out
}

func unfold(line string, counts []int) (string, []int) {
	var outCounts []int
	for i := 0; i < len(counts)*5; i++ {
		outCounts = append(outCounts, counts[i%len(counts)])
	}

	return line + "?" + line + "?" + line + "?" + line + "?" + line, outCounts
}
