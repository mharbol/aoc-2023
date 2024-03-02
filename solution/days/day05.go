package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mharbol/aoc-2023/util/span"
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

	seedSpans, orderedMappers := parsePart2(lines)

	for level, mapper := range orderedMappers {
		seedSpans = mapper.getDesitnationsForAllSpans(seedSpans, level)
	}

	minVal := 1<<31 - 1 // max int32 val
	for _, endSpan := range seedSpans {
		if endSpan.Start() < minVal {
			minVal = endSpan.Start()
		}
	}

	return fmt.Sprint(minVal), nil
}

// basic struct to determine if an `int` is in a certain span and map to an output value
// or span of outputs
type spanMapper struct {
	sourceSpan  *span.Span
	outputStart int
}

func newSpanMapper(sourceStart, outputStart, givenSpan int) *spanMapper {
	return &spanMapper{
		sourceSpan:  span.New(sourceStart, sourceStart+givenSpan),
		outputStart: outputStart,
	}
}

func newSpanMapperFromString(str string) *spanMapper {
	numsStrSlc := strings.Split(str, " ")
	desStart, _ := strconv.Atoi(numsStrSlc[0])
	srcStart, _ := strconv.Atoi(numsStrSlc[1])
	initSpan, _ := strconv.Atoi(numsStrSlc[2])
	return newSpanMapper(srcStart, desStart, initSpan)
}

// checks to see if the given source INT!!! is inside the span for this
// spanChecker's definition
func (sm *spanMapper) isIntInSpanMapper(source int) bool {
	return sm.sourceSpan.IsIntInSpan(source)
}

// returns the output for the given input int.
// This assumes that the input is in the defined spanMapper
func (sm *spanMapper) getSpannedOutputInt(input int) int {
	delta := input - sm.sourceSpan.Start()
	return sm.outputStart + delta
}

// returns the outputs for the given input Span
// first is the mapped output span which this spanMapper maps to
// second is the Span which "fell through" (does not map to anything prior), this is the left side of the overlap
// third is the remainder, this is what other spanMapper(s) will deal with, this is the right side of the overlap
// any empty returns are nil
func (sm *spanMapper) getSpannedOutputsForSpan(input *span.Span) (*span.Span, *span.Span, *span.Span) {
	var overlapSpan *span.Span = nil
	var leftSpan *span.Span = nil
	var rightSpan *span.Span = nil

	// get the [potential] overlap
	common, _, other := sm.sourceSpan.Venn(input)

	// map common
	if common != nil {
		deltaStart := common.Start() - sm.sourceSpan.Start()
		deltaEnd := common.End() - sm.sourceSpan.Start()

		overlapSpan = span.New(sm.outputStart+deltaStart, sm.outputStart+deltaEnd)
	}

	// get the left and right outliers
	for _, outlier := range other {
		// check for left
		if outlier.End() <= sm.sourceSpan.Start() {
			leftSpan = outlier
		}
		// check for right
		if outlier.Start() <= sm.sourceSpan.End() {
			rightSpan = outlier
		}
	}

	return overlapSpan, leftSpan, rightSpan
}

// compares spanCheckers for the purpose of sorting
// sorting criteria is value of sourceStart
func (sm *spanMapper) compareTo(other *spanMapper) int {
	return sm.sourceSpan.CompareTo(other.sourceSpan)
}

// aggregate of spanCheckers to determine the destination of any given input value
type resourceMapper struct {
	spanCheckers []*spanMapper
}

func newResourceMapperFromStringSlice(strings []string) *resourceMapper {
	var checkers []*spanMapper

	for _, checkerStr := range strings {
		checkers = append(checkers, newSpanMapperFromString(checkerStr))
	}

	return &resourceMapper{spanCheckers: checkers}
}

func (rm *resourceMapper) getDesination(input int) int {
	for _, mapper := range rm.spanCheckers {
		if mapper.isIntInSpanMapper(input) {
			return mapper.getSpannedOutputInt(input)
		}
	}
	// if none found, use the input number as the destination
	return input
}

func (rm *resourceMapper) sort() {
	sort.Slice(rm.spanCheckers, func(i, j int) bool {
		return rm.spanCheckers[i].compareTo(rm.spanCheckers[j]) < 0
	})
}

// ASSUMES SORTED RESOURCE MAPPER
// walks s through the spanMappers for this resourceMapper.
// overlaps and leftSpans that were not cought by the overlap get added the slice that refers to the next level of resources
// the right span will then overwrite s and be the new span in question
// this continues to the end or if s evenually becoems nil
func (rm *resourceMapper) getDesitnationsForSpan(s *span.Span) []*span.Span {
	var out []*span.Span

	sp := span.New(s.Start(), s.End())

	for _, checker := range rm.spanCheckers {
		mapped, leftSpan, rightSpan := checker.getSpannedOutputsForSpan(sp)

		if mapped != nil {
			out = append(out, mapped)
			if leftSpan != nil {
				out = append(out, leftSpan)
			}
			if rightSpan != nil {
				sp = rightSpan
			} else {
				sp = nil
				break
			}
		}
	}
	if sp != nil {
		out = append(out, sp)
	}
	return out
}

func (rm *resourceMapper) getDesitnationsForAllSpans(spans []*span.Span, level int) []*span.Span {

	var out []*span.Span

	for _, sp := range spans {
		out = append(out, rm.getDesitnationsForSpan(sp)...)
	}
	return span.Combine(out)
}

// Functions for parsing input to this problem //

func parsePart1(lines []string) ([]int, []*resourceMapper) {
	seeds := parseSeedsPart1(lines[0])
	entries := getBlocksOfEntries(lines)
	var mappers []*resourceMapper
	for _, entry := range entries {
		mappers = append(mappers, newResourceMapperFromStringSlice(entry))
	}

	return seeds, mappers
}

func parsePart2(lines []string) ([]*span.Span, []*resourceMapper) {
	seedSpans := parseSeedsPart2(lines[0])
	entries := getBlocksOfEntries(lines)
	var mappers []*resourceMapper
	for _, entry := range entries {
		mapper := newResourceMapperFromStringSlice(entry)
		mapper.sort()
		mappers = append(mappers, mapper)
	}
	return seedSpans, mappers
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

func parseSeedsPart2(line string) []*span.Span {
	var ret []*span.Span

	seedsStrArr := strings.Split(line[7:], " ")
	for idx := 0; idx < len(seedsStrArr); idx += 2 {
		start, _ := strconv.Atoi(seedsStrArr[idx])
		length, _ := strconv.Atoi(seedsStrArr[idx+1])
		ret = append(ret, span.New(start, start+length))
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
