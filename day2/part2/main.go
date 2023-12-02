package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"strings"
)

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

func main() {
	input, err := util.GetPart2InputFileLines("./day2")
	//input, err := util.GetExample2InputFileLines("./day2")
	util.Check(err)

	sumOfGames := 0
	for i := 0; i < len(input); i++ {
		sumOfGames += minimumCubePower(input[i])
	}

	fmt.Printf("Part 2 answer: %d", sumOfGames)
}

func minimumCubePower(gameInput string) int {
	maxRed, maxGreen, maxBlue := 0, 0, 0

	allPulls := strings.Split(gameInput, ":")
	individualPulls := strings.Split(allPulls[1], ";")
	//  3 blue, 4 red
	//  1 red, 2 green 6 blue
	// 2 green

	for i := 0; i < len(individualPulls); i++ {
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
				if colorNum > maxRed {
					maxRed = colorNum
				}
				break
			case green:
				colorNum, err := strconv.Atoi(numAndColors[1])
				util.Check(err)
				if colorNum > maxGreen {
					maxGreen = colorNum
				}
				break
			case blue:
				colorNum, err := strconv.Atoi(numAndColors[1])
				util.Check(err)
				if colorNum > maxBlue {
					maxBlue = colorNum
				}
				break
			default:
				err := "Invalid color as input: " + numAndColors[2]
				panic(err)
			}
		}
	}
	return maxRed * maxGreen * maxBlue
}
