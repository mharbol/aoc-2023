package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day05Part1(lines []string) (string, error) {

	var vals []int

	seeds, mappers := parsePart1(lines)

	for _, seedNum := range seeds {
		val := seedNum
		for _, mapper := range mappers {
			val = mapper.getDesination(val)
		}
		vals = append(vals, val)
	}

	sort.Ints(vals)

	return fmt.Sprint(vals[0]), nil
}

func Day05Part2(lines []string) (string, error) {

	_, mappers := parsePart1(lines)
	seedPairs := part2Seeds(lines)
	var vals []int

	for _, seedPair := range seedPairs {
		for seedPair.current() != -1 {
			num := seedPair.getNextSeedNumber()
			if num == -1 {
				continue
			}
			val := num
			for _, mapper := range mappers {
				val = mapper.getDesination(val)
			}
			vals = append(vals, val)
		}
	}

	sort.Ints(vals)

	return fmt.Sprint(vals[0]), nil
}

type spanChecker struct {
	sourceStart int
	sourceEnd   int
	outputStart int
}

// checks to see if the given source is inside the range for this
// spanChecker's definition
func (sc *spanChecker) isInRange(source int) bool {
	return source >= sc.sourceStart && source < sc.sourceEnd // strictly less than because the span starts on sourceStart
}

// returns the output for the given input.
// This assumes that the input is in the defined range
func (sc *spanChecker) getOutput(input int) int {
	delta := input - sc.sourceStart
	return sc.outputStart + delta
}

func newSpanChecker(sourceStart, outputStart, givenSpan int) *spanChecker {
	return &spanChecker{
		sourceStart: sourceStart,
		outputStart: outputStart,
		sourceEnd:   sourceStart + givenSpan,
	}
}

func makeSpanCheckerFromString(str string) *spanChecker {
	numsStrSlc := strings.Split(str, " ")
	desStart, _ := strconv.Atoi(numsStrSlc[0])
	srcStart, _ := strconv.Atoi(numsStrSlc[1])
	initSpan, _ := strconv.Atoi(numsStrSlc[2])
	return newSpanChecker(srcStart, desStart, initSpan)
}

type resourceMapper struct {
	spanCheckers []*spanChecker
}

func (rm *resourceMapper) getDesination(input int) int {
	// loop throught the spanCheckers to see if any are in range
	// return the output if so
	for _, checker := range rm.spanCheckers {
		if checker.isInRange(input) {
			return checker.getOutput(input)
		}
	}
	// if none found, use the input number as the destination
	return input
}

func newResourceMapper(strings []string) *resourceMapper {
	var checkers []*spanChecker

	for _, checkerStr := range strings {
		checkers = append(checkers, makeSpanCheckerFromString(checkerStr))
	}

	return &resourceMapper{spanCheckers: checkers}
}

func parseSeeds(line string) []int {
	var ret []int
	seedsStrArr := strings.Split(line[7:], " ")
	for _, numStr := range seedsStrArr {
		num, _ := strconv.Atoi(numStr)
		ret = append(ret, num)
	}
	return ret
}

func parsePart1(lines []string) ([]int, []*resourceMapper) {
	seeds := parseSeeds(lines[0])
	entries := getBlocksOfEntries(lines)
	var mappers []*resourceMapper
	for _, entry := range entries {
		mappers = append(mappers, newResourceMapper(entry))
	}

	return seeds, mappers
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

type p2seedPair struct {
	seedNum int
	span    int
	count   int
}

func (sp *p2seedPair) getNextSeedNumber() int {
	if sp.count == sp.span {
		return -1
	}
	sp.count++
	fmt.Println("Seed is", sp.seedNum, "span is", sp.span, "count is", sp.count)
	return sp.seedNum + sp.count
}

func (sp *p2seedPair) current() int {
    if sp.count == sp.span {
        return -1
    }
    return 0
}

func part2Seeds(lines []string) []p2seedPair {
	var out []p2seedPair
	origSeeds := parseSeeds(lines[0])
	for index := 0; index < len(origSeeds); index += 2 {
		fmt.Println("adding", origSeeds[index], origSeeds[index+1])
		out = append(out, p2seedPair{origSeeds[index], origSeeds[index+1], 0})
	}
	return out
}
