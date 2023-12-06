package main

import (
	"adventOfCode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// day6 Part 2
func main() {
	// Input as string

	input, err := util.GetPart2InputFileAsString("./day6")
	//input, err := util.GetExample2InputFileAsString("./day6")
	util.Check(err)

	lines := regexp.MustCompile("\r?\n").Split(input, -1)
	timesString := strings.Split(lines[0], "Time:")[1]
	distancesString := strings.Split(lines[1], "Distance:")[1]

	time, distance := parseNumberString(timesString), parseNumberString(distancesString)

	sumOfWins := 0
	for j := 0; j < time; j++ {
		distanceCovered := j * (time - j)
		if distanceCovered > distance {
			sumOfWins++
		}
	}

	fmt.Printf("The answer for part2 is: %d", sumOfWins)
}

func parseNumberString(numberString string) int {
	runes := []rune(numberString)

	var s string
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			s += string(runes[i])
		}
	}
	number, _ := strconv.Atoi(s)
	return number
}
