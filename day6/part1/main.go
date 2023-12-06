package main

import (
	"adventOfCode2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// day6 Part 1
func main() {
	// Input as string

	input, err := util.GetPart1InputFileAsString("./day6")
	//input, err := util.GetExample1InputFileAsString("./day6")
	util.Check(err)

	lines := regexp.MustCompile("\r?\n").Split(input, -1)
	timesString := strings.Split(lines[0], "Time:")[1]
	distancesString := strings.Split(lines[1], "Distance:")[1]

	times, distances := parseNumberString(timesString), parseNumberString(distancesString)

	answer := 1
	sumOfWins := 0
	for i := 0; i < len(times); i++ {
		for j := 0; j < times[i]; j++ {
			distance := j * (times[i] - j)
			if distance > distances[i] {
				sumOfWins++
			}
		}
		answer *= sumOfWins
		sumOfWins = 0
	}

	fmt.Printf("The answer for part1 is: %d", answer)
}

func parseNumberString(numberString string) []int {

	var numArray []int
	runes := []rune(numberString)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			s := string(runes[i])

			for j := i + 1; j < len(runes); j++ {
				if unicode.IsDigit(runes[j]) {
					temp := string(runes[j])
					s += temp
				} else {
					break
				}
				i = j + 1
			}
			num, _ := strconv.Atoi(s)
			numArray = append(numArray, num)
		}
	}

	return numArray
}
