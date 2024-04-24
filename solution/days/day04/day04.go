package day04

import (
	"fmt"
	"regexp"
	"strings"
)

func Part1(lines []string) string {

	acc := 0

	for _, line := range lines {
		numWins := countWins(line)
		acc += scoreWins(numWins)
	}

	return fmt.Sprint(acc)
}

func Part2(lines []string) string {

	acc := 0
	// defaults to all 0's
	countCards := make([]int, len(lines))

	for gameNumber, line := range lines {

		// first increase the number of cards by 1 for the original
		countCards[gameNumber] += 1

		// number of new tickets to give out
		numMatches := countWins(line)

		// loop down the slice numMatches times adding the number of
		// cards we had for this game
		numCardsPerGame := countCards[gameNumber]

		for index := gameNumber + 1; index < len(lines) && index < gameNumber+numMatches+1; index++ {
			countCards[index] += numCardsPerGame
		}
		acc += numCardsPerGame
	}
	return fmt.Sprint(acc)
}

func countWins(line string) int {

	// get the winning and user numbers as slices
	cardBody := line[strings.IndexByte(line, ':')+1:]
	cardBody = strings.TrimSpace(cardBody)

	leftAndRight := strings.Split(cardBody, "|")

	winningNumbers := strings.TrimSpace(leftAndRight[0])
	playerNumbers := strings.TrimSpace(leftAndRight[1])

	re := regexp.MustCompile(" +")

	// split over the spaces
	winningNumbersSlice := re.Split(winningNumbers, -1)
	playerNumbersSlice := re.Split(playerNumbers, -1)

	// place the winning numbers in a map where the
	// number is the key and the value is the count
	// (initially 0)
	winningNumbersCount := make(map[string]int)
	for _, s := range winningNumbersSlice {
		winningNumbersCount[s] = 0
	}

	// add player numbers to the counts if in the map
	for _, s := range playerNumbersSlice {

		count, ok := winningNumbersCount[s]
		if ok {
			winningNumbersCount[s] = count + 1
		}
	}

	// sum up counts
	total := 0
	for _, value := range winningNumbersCount {
		total += value
	}

	return total
}

func scoreWins(numWins int) int {
	var score int

	if numWins == 0 {
		score = 0
	} else {
		score = 1
		numWins--
	}
	for numWins > 0 {
		score *= 2
		numWins--
	}
	return score
}
