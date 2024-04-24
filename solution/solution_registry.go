package solution

import (
	"github.com/mharbol/aoc-2023/solution/days/day01"
	"github.com/mharbol/aoc-2023/solution/days/day02"
	"github.com/mharbol/aoc-2023/solution/days/day03"
	"github.com/mharbol/aoc-2023/solution/days/day04"
	"github.com/mharbol/aoc-2023/solution/days/day05"
	"github.com/mharbol/aoc-2023/solution/days/day06"
	"github.com/mharbol/aoc-2023/solution/days/day07"
	"github.com/mharbol/aoc-2023/solution/days/day08"
	"github.com/mharbol/aoc-2023/solution/days/day09"
	"github.com/mharbol/aoc-2023/solution/days/day10"
	"github.com/mharbol/aoc-2023/solution/days/day11"
	"github.com/mharbol/aoc-2023/solution/days/day12"
	"github.com/mharbol/aoc-2023/solution/days/day13"
	"github.com/mharbol/aoc-2023/solution/days/day14"
	"github.com/mharbol/aoc-2023/solution/days/day15"
	"github.com/mharbol/aoc-2023/solution/days/day16"
	"github.com/mharbol/aoc-2023/solution/days/day17"
	"github.com/mharbol/aoc-2023/solution/days/day18"
	"github.com/mharbol/aoc-2023/solution/days/day19"
)

var allSolutions = map[uint8]*solFuncTuple{
	1:  {day01.Part1, day01.Part2},
	2:  {day02.Part1, day02.Part2},
	3:  {day03.Part1, day03.Part2},
	4:  {day04.Part1, day04.Part2},
	5:  {day05.Part1, day05.Part2},
	6:  {day06.Part1, day06.Part2},
	7:  {day07.Part1, day07.Part2},
	8:  {day08.Part1, day08.Part2},
	9:  {day09.Part1, day09.Part2},
	10: {day10.Part1, day10.Part2},
	11: {day11.Part1, day11.Part2},
	12: {day12.Part1, day12.Part2},
	13: {day13.Part1, day13.Part2},
	14: {day14.Part1, day14.Part2},
	15: {day15.Part1, day15.Part2},
	16: {day16.Part1, day16.Part2},
	17: {day17.Part1, day17.Part2},
	18: {day18.Part1, day18.Part2},
	19: {day19.Part1, day19.Part2},
}
