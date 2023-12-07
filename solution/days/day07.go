package days

import (
	"fmt"
	"sort"
	"strconv"
)

func Day07Part1(lines []string) (string, error) {

	var hands []camelCardHand

	for _, line := range lines {
		wager, _ := strconv.Atoi(line[6:])
		hands = append(hands, camelCardHand{line[:5], wager})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareHands(hands[j]) < 0
	})

	tot := 0

	for i, card := range hands {
		tot += card.bid * (i + 1)
	}

	return fmt.Sprint(tot), nil
}

func Day07Part2(lines []string) (string, error) {

	return "", nil
}

type camelCardHand struct {
	hand string
	bid  int
}

func (s camelCardHand) compareHands(o camelCardHand) int {
	sType := classifyHandP1(s.hand)
	oType := classifyHandP1(o.hand)
	if sType != oType {
		return sType - oType
	}
	for idx := 0; idx < 5; idx++ {
		if s.hand[idx] == o.hand[idx] {
			continue
		}
		return cardRankingP1[byte(s.hand[idx])] - cardRankingP1[byte(o.hand[idx])]
	}
	return 0
}

const (
	HIGH            = 0
	PAIR            = 1
	TWO_PAIR        = 2
	THREE_OF_A_KIND = 3
	FULL_HOUSE      = 4
	FOUR_OF_A_KIND  = 5
	FIVE_OF_A_KIND  = 6
)

func classifyHandP1(hand string) int {
	cardCount := map[byte]int{
		'A': 0,
		'K': 0,
		'Q': 0,
		'J': 0,
		'T': 0,
		'9': 0,
		'8': 0,
		'7': 0,
		'6': 0,
		'5': 0,
		'4': 0,
		'3': 0,
		'2': 0,
	}
	for _, symbol := range hand {
		cardCount[byte(symbol)]++
	}

	countTriples := 0
	countPairs := 0

	for _, count := range cardCount {
		if count == 0 {
			continue
		} else if count == 5 {
			return FIVE_OF_A_KIND
		} else if count == 4 {
			return FOUR_OF_A_KIND
		} else if count == 2 {
			countPairs++
		} else if count == 3 {
			countTriples++
		}
	}
	if countTriples == 1 {
		if countPairs == 1 {
			return FULL_HOUSE
		} else {
			return THREE_OF_A_KIND
		}
	}
	if countPairs == 2 {
		return TWO_PAIR
	}
	if countPairs == 1 {
		return PAIR
	}
	return HIGH
}

var cardRankingP1 = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}
