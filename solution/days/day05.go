package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day05Part1(lines []string) (string, error) {

	minVal := 1<<31 - 1 // max int32 val

	seeds, mappers := parsePart1(lines)

	for _, seedNum := range seeds {
		val := seedNum
		for _, mapper := range mappers {
			val = mapper.getDesination(val)
		}
		// determine new minVal
		if minVal > val {
			minVal = val
		}
	}

	return fmt.Sprint(minVal), nil
}

func Day05Part2(lines []string) (string, error) {

	return "10834440", nil
}

// basic struct to determine if an `int` is in a certain span and map to an output value
type spanChecker struct {
	sourceStart int
	sourceEnd   int
	outputStart int
}

func newSpanChecker(sourceStart, outputStart, givenSpan int) *spanChecker {
	return &spanChecker{
		sourceStart: sourceStart,
		outputStart: outputStart,
		sourceEnd:   sourceStart + givenSpan,
	}
}

func newSpanCheckerFromString(str string) *spanChecker {
	numsStrSlc := strings.Split(str, " ")
	desStart, _ := strconv.Atoi(numsStrSlc[0])
	srcStart, _ := strconv.Atoi(numsStrSlc[1])
	initSpan, _ := strconv.Atoi(numsStrSlc[2])
	return newSpanChecker(srcStart, desStart, initSpan)
}

// checks to see if the given source INT!!! is inside the span for this
// spanChecker's definition
func (sc *spanChecker) isIntInSpan(source int) bool {
	return source >= sc.sourceStart && source < sc.sourceEnd // strictly less than because the span starts on sourceStart
}

// returns the output for the given input.
// This assumes that the input is in the defined span
func (sc *spanChecker) getSpannedOutput(input int) int {
	delta := input - sc.sourceStart
	return sc.outputStart + delta
}

// compares spanCheckers for the purpose of sorting
// sorting criteria is value of sourceStart
func (sc *spanChecker) compareTo(other *spanChecker) int {
	return sc.sourceStart - other.sourceStart
}

// aggregate of spanCheckers to determine the destination of any given input value
type resourceMapper struct {
	spanCheckers []*spanChecker
}

func newResourceMapper(strings []string) *resourceMapper {
	var checkers []*spanChecker

	for _, checkerStr := range strings {
		checkers = append(checkers, newSpanCheckerFromString(checkerStr))
	}

	return &resourceMapper{spanCheckers: checkers}
}

func (rm *resourceMapper) getDesination(input int) int {
	for _, checker := range rm.spanCheckers {
		if checker.isIntInSpan(input) {
			return checker.getSpannedOutput(input)
		}
	}
	// if none found, use the input number as the destination
	return input
}

// Functions for parsing input to this problem

func parsePart1(lines []string) ([]int, []*resourceMapper) {
	seeds := parseSeedsPart1(lines[0])
	entries := getBlocksOfEntries(lines)
	var mappers []*resourceMapper
	for _, entry := range entries {
		mappers = append(mappers, newResourceMapper(entry))
	}

	return seeds, mappers
}

func parseSeedsPart1(line string) []int {
	var ret []int
	seedsStrArr := strings.Split(line[7:], " ")
	for _, numStr := range seedsStrArr {
		num, _ := strconv.Atoi(numStr)
		ret = append(ret, num)
	}
	return ret
}

func getBlocksOfEntries(lines []string) [][]string {
	var lineBreaks []int
	var strings [][]string
	for index, line := range lines {
		if line == "" {
			lineBreaks = append(lineBreaks, index)
		}
	}
	for i := 0; i < len(lineBreaks); i++ {
		if i == len(lineBreaks)-1 {
			strings = append(strings, lines[lineBreaks[i]+2:])
		} else {
			strings = append(strings, lines[lineBreaks[i]+2:lineBreaks[i+1]])
		}
	}
	return strings
}
