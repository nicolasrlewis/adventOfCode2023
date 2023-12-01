package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	input, err := util.GetPart2InputFileLines("./day1")
	//input, err := util.GetExample2InputFileLines("./day1")
	util.Check(err)

	// Part 2 strings
	digitStringMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0
	var first, second rune
	firstFound, secondFound := false, false
	var calibrationValue int
	var codeString string
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
			} else {
				codeString = string(r)
				for k := j + 1; k < len(runes); k++ {
					r2 := runes[k]
					if unicode.IsDigit(r2) {
						break
					}
					codeString += string(r2)

					val, found := digitStringMap[codeString]

					if found {
						if firstFound {
							second = rune('0' + val)
							secondFound = true
						} else {
							first = rune('0' + val)
							firstFound = true
						}
						break
					}
				}
			}
		}
		if !secondFound {
			second = first
		}
		calibrationValue, err = strconv.Atoi(string(first) + string(second))
		util.Check(err)
		fmt.Println(calibrationValue)
		sum += calibrationValue

		calibrationValue = 0
		firstFound = false
		secondFound = false
		codeString = ""
	}
	fmt.Println(sum)
}
