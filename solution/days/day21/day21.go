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

	total := 0
	for _, count := range walkMap {
		if count%2 == 0 {
			total++
		}
	}

	return fmt.Sprint(total)
}

func Part2(lines []string) string {

	return ""
}
