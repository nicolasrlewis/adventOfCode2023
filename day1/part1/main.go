package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	input, err := util.GetPart1InputFileLines("./day1")
	//input, err := util.GetExample1InputFileLines("./day1")
	util.Check(err)

	sum := 0
	var first, second rune
	firstFound, secondFound := false, false
	var calibrationValue int
	for i := 0; i < len(input); i++ {
		runes := []rune(input[i])
		for j := 0; j < len(runes); j++ {
			r := runes[j]
			if unicode.IsDigit(r) {
				if firstFound {
					second = r
					secondFound = true
				} else {
					first = r
					firstFound = true
				}
			}
		}
		if !secondFound {
			second = first
		}
		calibrationValue, err = strconv.Atoi(string(first) + string(second))
		util.Check(err)
		sum += calibrationValue

		calibrationValue = 0
		firstFound = false
		secondFound = false
	}
	fmt.Println(sum)
}
