package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"strings"
)

const (
	red                       = "red"
	green                     = "green"
	blue                      = "blue"
	maxRed, maxGreen, maxBlue = 12, 13, 14
)

// Day2 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day2")
	//input, err := util.GetExample1InputFileLines("./day2")
	util.Check(err)

	sumOfGames := 0
	for i, inputLine := range input {
		if isGamePossible(inputLine) {
			sumOfGames += i + 1
		}
	}

	fmt.Printf("Part 1 answer: %d", sumOfGames)
}

func isGamePossible(gameInput string) bool {
	allPulls := strings.Split(gameInput, ":")
	individualPulls := strings.Split(allPulls[1], ";")
	//  3 blue, 4 red
	//  1 red, 2 green 6 blue
	// 2 green

	for i := 0; i < len(individualPulls); i++ {
		numOfRed, numOfGreen, numOfBlue := 0, 0, 0

		colors := strings.Split(individualPulls[i], ",")
		// Note the leading spaces
		//  3 blue
		//  4 red
		for j := 0; j < len(colors); j++ {
			numAndColors := strings.Split(colors[j], " ")
			//
			// 3
			// blue
			switch numAndColors[2] {
			case red:
				colorNum, err := strconv.Atoi(numAndColors[1])
				util.Check(err)
				numOfRed += colorNum
				break
			case green:
				colorNum, err := strconv.Atoi(numAndColors[1])
				util.Check(err)
				numOfGreen += colorNum
				break
			case blue:
				colorNum, err := strconv.Atoi(numAndColors[1])
				util.Check(err)
				numOfBlue += colorNum
				break
			default:
				err := "Invalid color as input: " + numAndColors[2]
				panic(err)
			}
		}

		if numOfRed > maxRed || numOfGreen > maxGreen || numOfBlue > maxBlue {
			return false
		}
	}
	return true
}
