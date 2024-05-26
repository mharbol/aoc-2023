package day21

import "fmt"

func Part1(lines []string) string {

	const MAX_STEP int = 64

	walkMap := make(map[coord]int)

	walkMap[findStart(lines)] = 0

	for lastStep := 0; lastStep < MAX_STEP; lastStep++ {
		for coordinate, count := range walkMap {
			if count == lastStep {
				walkFromPoint(coordinate, lines, walkMap, lastStep+1)
			}
		}
	}

	return fmt.Sprint(countEven(walkMap))
}

func Part2(lines []string) string {

    return ""
}

func countEven(walkMap map[coord]int) int {
	total := 0
	for _, count := range walkMap {
		if count%2 == 0 {
			total++
		}
	}
	return total
}

func countOdd(walkMap map[coord]int) int {
	total := 0
	for _, count := range walkMap {
		if count%2 == 1 {
			total++
		}
	}
	return total
}
