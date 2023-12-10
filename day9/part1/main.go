package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"strings"
)

// day9 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day9")
	//input, err := util.GetExample1InputFileLines("./day9")
	util.Check(err)

	sum := 0
	for i := 0; i < len(input); i++ {
		splitInput := strings.Split(input[i], " ")

		numbers := make([]int, len(splitInput))
		for j := 0; j < len(splitInput); j++ {
			numbers[j], _ = strconv.Atoi(splitInput[j])
		}

		var sequences [][]int
		sequences = append(sequences, numbers)
		for j := 0; j < len(sequences); j++ {
			differences := make([]int, len(sequences[j])-1)
			allZeros := true
			for k := 0; k < len(sequences[j])-1; k++ {
				difference := sequences[j][k+1] - sequences[j][k]
				differences[k] = difference
				if difference != 0 {
					allZeros = false
				}
			}
			sequences = append(sequences, differences)
			if allZeros {
				break
			}
		}
		seqIndex := len(sequences) - 1
		sequences[seqIndex] = append(sequences[seqIndex], sequences[seqIndex][len(sequences[seqIndex])-1])
		for j := len(sequences) - 2; j >= 0; j-- {
			prevSubArrayIndex := len(sequences[j+1]) - 1
			subArrayIndex := len(sequences[j]) - 1
			num := sequences[j+1][prevSubArrayIndex] + sequences[j][subArrayIndex]
			sequences[j] = append(sequences[j], num)

			if j == 0 {
				sum += num
			}
		}
	}

	fmt.Printf("The answer to part1 is: %d", sum)
}
