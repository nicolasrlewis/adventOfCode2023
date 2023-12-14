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

const expansionNum = 1000000

// day11 Part 2
func main() {
	input, err := util.GetPart2InputFileLines("./day11")
	//input, err := util.GetExample2InputFileLines("./day11")
	util.Check(err)

	pairs := buildPairPoints(input)
	distances := findDistances(pairs)

	sum := 0
	for i := 0; i < len(distances); i++ {
		sum += distances[i]
	}

	fmt.Printf("The answer for part1 is: %d", sum)

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
	numOfEmptyRows, numOfEmptyCols := 0, 0

	for i := 0; i < len(input); i++ {
		if isEmptyRow(input[i]) {
			numOfEmptyRows++
		}
		for j := 0; j < len(input[i]); j++ {
			if string(input[i][j]) == "#" {
				xScalar, yScalar := 0, 0
				if numOfEmptyRows != 0 {
					xScalar = (expansionNum - 1) * numOfEmptyRows
				}
				if numOfEmptyCols != 0 {
					yScalar = (expansionNum - 1) * numOfEmptyCols
				}
				points = append(points, Point{x: i + xScalar, y: j + yScalar})

				numOfEmptyCols = 0
			} else if isEmptyColumn(input, j) {
				numOfEmptyCols++
			}
		}
		numOfEmptyCols = 0
	}
	return points
}

func isEmptyRow(row string) bool {
	for i := 0; i < len(row); i++ {
		if string(row[i]) != "." {
			return false
		}
	}
	return true
}

func isEmptyColumn(col []string, index int) bool {
	for i := 0; i < len(col); i++ {
		if string(col[i][index]) != "." {
			return false
		}
	}
	return true
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
