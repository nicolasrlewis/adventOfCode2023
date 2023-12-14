package main

import (
	"adventOfCode2023/util"
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type PointPair struct {
	point1 Point
	point2 Point
}

// day11 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day11")
	//input, err := util.GetExample1InputFileLines("./day11")
	util.Check(err)

	input = expand(input)

	pairs := buildPairPoints(input)
	distances := findDistances(pairs)

	sum := 0
	for i := 0; i < len(distances); i++ {
		sum += distances[i]
	}

	fmt.Printf("The answer for part1 is: %d", sum)

}

func expand(input []string) []string {
	input = expandRows(input)
	input = expandColumns(input)

	return input
}

func expandRows(input []string) []string {
	foundGalaxy := false
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if string(input[i][j]) == "#" {
				foundGalaxy = true
				break
			}
		}
		if !foundGalaxy {
			// expand
			input = addEmptySpaceRow(input, i)
			i++
		}
		foundGalaxy = false
	}
	return input
}

func expandColumns(input []string) []string {

	foundGalaxy := false
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			if string(input[j][i]) == "#" {
				foundGalaxy = true
				break
			}
		}
		if !foundGalaxy {
			// expand
			input = addEmptySpaceColumn(input, i)
			i++
		}
		foundGalaxy = false
	}

	return input
}

func addEmptySpaceRow(input []string, index int) []string {
	emptySpaceSize := len(input[index])

	emptySpaceString := ""
	for i := 0; i < emptySpaceSize; i++ {
		emptySpaceString += "."
	}
	return util.Insert(input, index, emptySpaceString).([]string)
}

func addEmptySpaceColumn(input []string, index int) []string {
	for i := 0; i < len(input); i++ {
		input[i] = util.InsertRune(input[i], index, '.')
	}
	return input
}

func buildPairPoints(input []string) []PointPair {
	points := findPoints(input)

	var pairs []PointPair
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, PointPair{
				point1: points[i],
				point2: points[j],
			})
		}
	}

	return pairs
}

func findPoints(input []string) []Point {
	var points []Point
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if string(input[i][j]) == "#" {
				points = append(points, Point{x: i, y: j})
			}
		}
	}
	return points
}

func findDistances(pairs []PointPair) []int {
	var distances []int
	for i := 0; i < len(pairs); i++ {
		distances = append(distances, calculateDistance(pairs[i].point1, pairs[i].point2))
	}
	return distances
}

func calculateDistance(point1 Point, point2 Point) int {
	return int(math.Abs(float64(point1.x-point2.x)) + math.Abs(float64(point1.y-point2.y)))
}
