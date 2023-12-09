package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardStrengths = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 0,
	"T": 10,
}

func GetTotalWinnings(documentName string) (totalWinnings int) {
	file, _ := os.Open(documentName)
	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		hands = append(hands, parseHand(scanner.Text()))
	}
	sort.Sort(ByScore(hands))
	fmt.Println(filter(hands))
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}
	return
}

func filter(ss []Hand) (ret []Hand) {
	for _, s := range ss {
		if strings.Contains(s.hand, "J") {
			ret = append(ret, s)
		}
	}
	return
}

func parseHand(line string) Hand {
	hand := strings.Fields(line)[0]
	bid, _ := strconv.Atoi(strings.Fields(line)[1])
	return Hand{
		hand:  hand,
		bid:   bid,
		score: getScore(hand),
	}
}

func getScore(hand string) int {
	if hasFiveOfAKind(hand) {
		return 6
	} else if hasFourOfAKind(hand) {
		return 5
	} else if hasFullHouse(hand) {
		return 4
	} else if hasThreeOfAKind(hand) {
		return 3
	} else if hasTwoPair(hand) {
		return 2
	} else if hasPair(hand) {
		return 1
	} else {
		return 0
	}
}

func hasFiveOfAKind(hand string) bool {
	cardMap := map[int32]int{'J': 0}
	for _, card := range hand {
		cardMap = initOrAdd(cardMap, card, 1)
		if card == 'J' {
			for key, num := range cardMap {
				if key != 'J' && num+cardMap['J'] == 5 {
					return true
				}
			}
		} else if cardMap[card]+cardMap['J'] == 5 {
			return true
		}
	}
	return cardMap['J'] == 5
}

func hasFourOfAKind(hand string) bool {
	cardMap := map[int32]int{'J': 0}
	for _, card := range hand {
		cardMap = initOrAdd(cardMap, card, 1)
		if card == 'J' {
			for key, num := range cardMap {
				if key != 'J' && num+cardMap['J'] >= 4 {
					return true
				}
			}
		} else if cardMap[card]+cardMap['J'] >= 4 {
			return true
		}
	}
	return false
}

func hasFullHouse(hand string) bool {
	cardMap := map[int32]int{'J': 0}
	for _, card := range hand {
		cardMap = initOrAdd(cardMap, card, 1)
		if card == 'J' {
			for key, num := range cardMap {
				if key != 'J' && num+cardMap['J'] > 3 {
					return false
				}
			}
		} else if cardMap[card]+cardMap['J'] > 3 {
			return false
		}
	}
	return len(cardMap) == 3
}

func hasThreeOfAKind(hand string) bool {
	cardMap := map[int32]int{'J': 0}
	for _, card := range hand {
		cardMap = initOrAdd(cardMap, card, 1)
		if card == 'J' {
			for key, num := range cardMap {
				if key != 'J' && num+cardMap['J'] >= 3 {
					return true
				}
			}
		} else if cardMap[card]+cardMap['J'] >= 3 {
			return true
		}
	}
	return false
}

func hasTwoPair(hand string) bool {
	cardMap := map[int32]int{'J': 0}
	numPairs := 0
	for _, card := range hand {
		cardMap = initOrAdd(cardMap, card, 1)
		if card == 'J' {
			numPairs += 1
			if numPairs == 2 {
				return true
			}
		} else if cardMap[card]%2 == 0 {
			numPairs += 1
			if numPairs == 2 {
				return true
			}
		}
	}
	return false
}

func hasPair(hand string) bool {
	cardMap := map[int32]int{'J': 0}
	for _, card := range hand {
		cardMap = initOrAdd(cardMap, card, 1)
		if card == 'J' {
			for key, num := range cardMap {
				if key != 'J' && num+cardMap['J'] == 2 {
					return true
				}
			}
		} else if cardMap[card]+cardMap['J'] == 2 {
			return true
		}
	}
	return false
}

func getCardStrength(card string) int {
	strength, err := strconv.Atoi(card)
	if err != nil {
		strength = cardStrengths[card]
	}
	return strength
}

type Hand struct {
	hand  string
	bid   int
	score int
}

type ByScore []Hand

func (hands ByScore) Len() int { return len(hands) }
func (hands ByScore) Less(i, j int) bool {
	if hands[i].score == hands[j].score {
		for k, card := range hands[i].hand {
			if int(card) != int(hands[j].hand[k]) {
				thisStrength := getCardStrength(string(card))
				thatStrength := getCardStrength(string(hands[j].hand[k]))
				return thisStrength < thatStrength
			}
		}
		return true
	}
	return hands[i].score < hands[j].score
}
func (hands ByScore) Swap(i, j int) { hands[i], hands[j] = hands[j], hands[i] }

func initOrAdd(m map[int32]int, key int32, value int) map[int32]int {
	_, exists := m[key]
	if exists {
		m[key] += value
	} else {
		m[key] = value
	}
	return m
}
