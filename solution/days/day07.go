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
		return hands[i].compareHandsP1(hands[j]) < 0
	})

	tot := 0

	for i, card := range hands {
		tot += card.bid * (i + 1)
	}

	return fmt.Sprint(tot), nil
}

func Day07Part2(lines []string) (string, error) {

	var hands []camelCardHand

    tot := 0

	for _, line := range lines {
		wager, _ := strconv.Atoi(line[6:])
		hands = append(hands, camelCardHand{line[:5], wager})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareHandsP2(hands[j]) < 0
	})

	for i, card := range hands {
		tot += card.bid * (i + 1)
	}

	return fmt.Sprint(tot), nil
}

type camelCardHand struct {
	hand string
	bid  int
}

func (s camelCardHand) compareHandsP1(o camelCardHand) int {
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

func (s camelCardHand) compareHandsP2(o camelCardHand) int {
	sType := classifyHandP2(s.hand)
	oType := classifyHandP2(o.hand)
	if sType != oType {
		return sType - oType
	}
	for idx := 0; idx < 5; idx++ {
		if s.hand[idx] == o.hand[idx] {
			continue
		}
		return cardRankingP2[byte(s.hand[idx])] - cardRankingP2[byte(o.hand[idx])]
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

func getCardCount(hand string) map[byte]int {
	cardCount := make(map[byte]int)
	for _, symbol := range hand {
		value, ok := cardCount[byte(symbol)]
		if !ok {
			cardCount[byte(symbol)] = 1
		} else {
			cardCount[byte(symbol)] = value + 1
		}
	}
	return cardCount
}

func classifyHandP2(hand string) int {

	cardCount := getCardCount(hand)

	countQuad := 0
	countTriples := 0
	countPairs := 0
	countJokers := 0

	for sym, count := range cardCount {

		// parse the hand
		if count == 5 {
			return FIVE_OF_A_KIND
		} else if sym == 'J' {
			countJokers = count
		} else if count == 2 {
			countPairs++
		} else if count == 3 {
			countTriples++
		} else if count == 4 {
			countQuad++
		}

	}

	// assess hand with no jokers
	if countJokers == 0 {
		if countQuad == 1 {
			return FOUR_OF_A_KIND
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

	// easy case, 4 jokers, with any card 5 of a kind
	if countJokers == 4 {
		return FIVE_OF_A_KIND
	}

	// three jokers, will make either 4 or 5 of a kind
	if countJokers == 3 {
		if countPairs == 1 {
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	}

	// two jokers
	if countJokers == 2 {
		if countTriples == 1 {
			return FIVE_OF_A_KIND
		}
		if countPairs == 1 {
			return FOUR_OF_A_KIND
		}
		return THREE_OF_A_KIND
	}

	// one joker
	if countJokers == 1 {
		if countQuad == 1 {
			return FIVE_OF_A_KIND
		}
		if countTriples == 1 {
			return FOUR_OF_A_KIND
		}
		if countPairs == 2 {
			return FULL_HOUSE
		}
		if countPairs == 1 {
			return THREE_OF_A_KIND // beats TWO_PAIR
		}
		return PAIR
	}
	return HIGH
}

func classifyHandP1(hand string) int {

	cardCount := getCardCount(hand)

	countTriples := 0
	countPairs := 0

	for _, count := range cardCount {
		if count == 5 {
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

var cardRankingP2 = map[byte]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}
