package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Card struct {
	winningNumbers []int
	pickedNumbers  []int
	copies         int
}

// day4 Part 2
func main() {
	input, err := util.GetPart2InputFileLines("./day4")
	//input, err := util.GetExample2InputFileLines("./day4")
	util.Check(err)

	var cards []Card
	for i := 0; i < len(input); i++ {
		splitText := "Card "
		splitText += string(strconv.Itoa(i+1)) + ": "
		card := strings.Split(input[i], splitText)
		card = strings.Split(card[1], " | ")

		winningNumbers := buildNumberSlice(card[0])
		pickedNumbers := buildNumberSlice(card[1])
		copies := 1

		cards = append(cards, Card{
			winningNumbers: winningNumbers,
			pickedNumbers:  pickedNumbers,
			copies:         copies,
		})
	}

	fmt.Printf("Part2 answer is: %d", processCards(cards))
}

func buildNumberSlice(s string) []int {
	var numbers []int

	runes := []rune(s)
	numberString := ""
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if unicode.IsDigit(r) {
			numberString += string(r)
		} else {
			if numberString != "" {
				number, err := strconv.Atoi(numberString)
				util.Check(err)
				numbers = append(numbers, number)
				numberString = ""
			}
		}
	}

	if numberString != "" {
		number, err := strconv.Atoi(numberString)
		util.Check(err)
		numbers = append(numbers, number)
		numberString = ""
	}

	return numbers
}

func processCards(cards []Card) int {
	for i := 0; i < len(cards); i++ {
		for j := 0; j < cards[i].copies; j++ {
			matches := 0
			for k := 0; k < len(cards[i].pickedNumbers); k++ {
				if contains(cards[i].winningNumbers, cards[i].pickedNumbers[k]) {
					matches++
				}
			}
			if matches != 0 {
				for k := 1; k <= matches; k++ {
					cards[i+k].copies++
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(cards); i++ {
		sum += cards[i].copies
	}

	return sum
}

func contains(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}

	return false
}
