package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"strings"
)

type HandType int
type Card int

const (
	HighCard HandType = iota // Start at 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

const (
	Two Card = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	J
	Q
	K
	A
)

type Hand struct {
	cards    []Card
	handType HandType
	rank     int
	bid      int
}

// day7 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day7")
	//input, err := util.GetExample1InputFileLines("./day7")
	util.Check(err)

	var hands [][]Hand

	for i := 0; i <= int(FiveOfAKind); i++ {
		var arr []Hand
		hands = append(hands, arr)
	}

	for i := 0; i < len(input); i++ {
		hand := buildHand(input[i])
		hands[hand.handType] = append(hands[hand.handType], hand)
	}

	for i := 0; i < len(hands); i++ {
		if len(hands[i]) > 0 {
			hands[i] = sortHandTypes(hands[i])
		}
	}

	calculateRank(hands)

	totalWinnings := 0
	for i := 0; i < len(hands); i++ {
		for j := 0; j < len(hands[i]); j++ {
			totalWinnings += hands[i][j].bid * hands[i][j].rank
		}
	}

	fmt.Printf("The answer for part1 is: %d", totalWinnings)
}

func buildHand(s string) Hand {
	var hand Hand

	handAndBid := strings.Split(s, " ")

	hand.cards = buildCards(handAndBid[0])
	hand.bid, _ = strconv.Atoi(handAndBid[1])
	hand.rank = -1

	hand.handType = determineHandType(hand.cards)

	return hand
}

func buildCards(s string) []Card {
	var cards []Card

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '2':
			cards = append(cards, Two)
			break
		case '3':
			cards = append(cards, Three)
			break
		case '4':
			cards = append(cards, Four)
			break
		case '5':
			cards = append(cards, Five)
			break
		case '6':
			cards = append(cards, Six)
			break
		case '7':
			cards = append(cards, Seven)
			break
		case '8':
			cards = append(cards, Eight)
			break
		case '9':
			cards = append(cards, Nine)
			break
		case 'T':
			cards = append(cards, T)
			break
		case 'J':
			cards = append(cards, J)
			break
		case 'Q':
			cards = append(cards, Q)
			break
		case 'K':
			cards = append(cards, K)
			break
		case 'A':
			cards = append(cards, A)
			break
		}
	}
	return cards
}

func determineHandType(cards []Card) HandType {
	slice := make([]int, int(A)+1)
	for i := 0; i < len(cards); i++ {
		card := cards[i]
		slice[card]++
	}

	if isFiveOfAKind(slice) {
		return FiveOfAKind
	}
	if isFourOfAKind(slice) {
		return FourOfAKind
	}
	if isFullHouse(slice) {
		return FullHouse
	}
	if isThreeOfAKind(slice) {
		return ThreeOfAKind
	}
	if isTwoPair(slice) {
		return TwoPair
	}
	if isOnePair(slice) {
		return OnePair
	}
	return HighCard
}

func isFiveOfAKind(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 3 {
			for j := 0; j < len(arr); j++ {
				if arr[j] == 2 {
					return true
				}
			}
			break
		}
	}
	return false
}

func isThreeOfAKind(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(arr []int) bool {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 2 {
			count++
		}
	}
	return count == 2
}

func isOnePair(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 2 {
			return true
		}
	}
	return false
}

func sortHandTypes(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if hands[i].cards[0] != hands[j].cards[0] {
				if hands[i].cards[0] > hands[j].cards[0] {
					hands[i], hands[j] = hands[j], hands[i]
				}
			} else if hands[i].cards[1] != hands[j].cards[1] {
				if hands[i].cards[1] > hands[j].cards[1] {
					hands[i], hands[j] = hands[j], hands[i]
				}
			} else if hands[i].cards[2] != hands[j].cards[2] {
				if hands[i].cards[2] > hands[j].cards[2] {
					hands[i], hands[j] = hands[j], hands[i]
				}
			} else if hands[i].cards[3] != hands[j].cards[3] {
				if hands[i].cards[3] > hands[j].cards[3] {
					hands[i], hands[j] = hands[j], hands[i]
				}
			} else if hands[i].cards[4] != hands[j].cards[4] {
				if hands[i].cards[4] > hands[j].cards[4] {
					hands[i], hands[j] = hands[j], hands[i]
				}
			}
		}
	}
	return hands
}

func calculateRank(hands [][]Hand) {
	rankCounter := 1
	for i := 0; i < len(hands); i++ {
		for j := 0; j < len(hands[i]); j++ {
			hands[i][j].rank = rankCounter
			rankCounter++
		}
	}
}
