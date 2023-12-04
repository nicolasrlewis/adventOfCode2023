package main

import (
	"adventOfCode2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// day4 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day4")
	//input, err := util.GetExample1InputFileLines("./day4")
	util.Check(err)

	sum := 0
	for i := 0; i < len(input); i++ {
		splitText := "Card "
		splitText += string(strconv.Itoa(i+1)) + ": "
		card := strings.Split(input[i], splitText)
		card = strings.Split(card[1], " | ")

		winningNumbers := buildNumberSlice(card[0])
		pickedNumbers := buildNumberSlice(card[1])
		matches := 0
		for j := 0; j < len(pickedNumbers); j++ {
			if contains(winningNumbers, pickedNumbers[j]) {
				matches++
			}
		}
		if matches != 0 {
			sum += int(math.Pow(2, float64(matches-1)))
		}
	}
	fmt.Printf("Part1 answer is: %d\n", sum)
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

func contains(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}

	return false
}
