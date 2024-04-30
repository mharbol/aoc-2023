package day20

import (
	"fmt"
	"slices"
)

func Part1(lines []string) string {

	mods := parseModules(lines)

	highCount, lowCount := 0, 0

	for i := 0; i < 1000; i++ {
		mods = pressButton(mods, &highCount, &lowCount)
	}

	return fmt.Sprint(highCount * lowCount)
}

func Part2(lines []string) string {

	mods := parseModules(lines)

	// get the rxRoot
	var rxRoot string
	for name, value := range mods {
		if slices.Index(value.modList(), "rx") > -1 {
			rxRoot = name
			break
		}
	}

	// get the "feeders" for rxRoot
	feeders := make([]string, 0)
	for name, value := range mods {
		if slices.Contains(value.modList(), rxRoot) {
			feeders = append(feeders, name)
		}
	}
	numFeeders := len(feeders)

	feederMap := make(map[string]int)
	count := 0

	for len(feederMap) != numFeeders {
		count++
		mods = pressButtonPart2(mods, rxRoot, feederMap, count)
	}

	prod := 1
	for _, num := range feederMap {
		prod *= num
	}

	return fmt.Sprint(prod)
}
