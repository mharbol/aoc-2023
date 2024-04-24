package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) string {

	acc := 0

	const (
		RED_MAX   int = 12
		GREEN_MAX int = 13
		BLUE_MAX  int = 14
	)

	for index, line := range lines {

		gameNum := index + 1

		// prune off game
		var start int
		if gameNum < 10 {
			start = 8
		} else if gameNum < 100 {
			start = 9
		} else {
			start = 10
		}
		game := line[start:]

		draws := strings.Split(game, "; ")

		isPossible := true

		for _, draw := range draws {
			pairs := strings.Split(draw, ", ")

			for _, pairStr := range pairs {
				var pair []string = strings.Split(pairStr, " ")
				count, _ := strconv.Atoi(pair[0])
				cubeColor := pair[1]
				switch cubeColor {
				case "red":
					if count > RED_MAX {
						isPossible = false
					}
					break
				case "green":
					if count > GREEN_MAX {
						isPossible = false
					}
					break
				case "blue":
					if count > BLUE_MAX {
						isPossible = false
					}
					break
				}
			}
		}
		if isPossible {
			acc += gameNum
		}
	}

	return fmt.Sprint(acc)
}

func Part2(lines []string) string {

	acc := 0

	for index, line := range lines {

		gameNum := index + 1

		// prune off game
		var start int
		if gameNum < 10 {
			start = 8
		} else if gameNum < 100 {
			start = 9
		} else {
			start = 10
		}
		game := line[start:]

		minRed := 0
		minBlue := 0
		minGreen := 0

		draws := strings.Split(game, "; ")

		for _, draw := range draws {
			pairs := strings.Split(draw, ", ")

			for _, pairStr := range pairs {
				var pair []string = strings.Split(pairStr, " ")
				count, _ := strconv.Atoi(pair[0])
				cubeColor := pair[1]

				switch cubeColor {
				case "red":
					if count > minRed {
						minRed = count
					}
					break
				case "green":
					if count > minGreen {
						minGreen = count
					}
					break
				case "blue":
					if count > minBlue {
						minBlue = count
					}
					break
				}
			}
		}
		power := minBlue * minRed * minGreen
		acc += power
	}

	return fmt.Sprint(acc)
}
