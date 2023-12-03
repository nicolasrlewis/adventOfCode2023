package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strconv"
	"unicode"
)

type Point struct {
	x, y int
}

// day3 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day3")
	//input, err := util.GetExample1InputFileLines("./day3")
	util.Check(err)

	var symbols []Point
	for i := 0; i < len(input); i++ {
		runes := []rune(input[i])
		for j := 0; j < len(runes); j++ {
			r := runes[j]
			if unicode.IsDigit(r) || r == '.' {
				continue
			}
			symbols = append(symbols, Point{x: i, y: j})
		}
	}

	sum := 0
	for i := 0; i < len(symbols); i++ {
		sum += calculateSymbolSum(input, symbols[i])
	}

	fmt.Printf("Sum of part numbers is: %d", sum)
}

func calculateSymbolSum(input []string, symbol Point) int {
	sum := 0
	var visited []Point
	if checkTopLeft(input, symbol) {
		var number int
		number, visited = grabNumber(input, Point{x: symbol.x - 1, y: symbol.y - 1}, visited)
		sum += number
	}

	if checkTop(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x - 1, y: symbol.y}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x - 1, y: symbol.y}, visited)
			sum += number
		}
	}

	if checkTopRight(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x - 1, y: symbol.y + 1}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x - 1, y: symbol.y + 1}, visited)
			sum += number
		}
	}

	if checkRight(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x, y: symbol.y + 1}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x, y: symbol.y + 1}, visited)
			sum += number
		}
	}

	if checkBottomRight(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x + 1, y: symbol.y + 1}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x + 1, y: symbol.y + 1}, visited)
			sum += number
		}
	}

	if checkBottom(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x + 1, y: symbol.y}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x + 1, y: symbol.y}, visited)
			sum += number
		}
	}

	if checkBottomLeft(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x + 1, y: symbol.y - 1}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x + 1, y: symbol.y - 1}, visited)
			sum += number
		}
	}

	if checkLeft(input, symbol) {
		if !checkVisited(visited, Point{x: symbol.x, y: symbol.y - 1}) {
			var number int
			number, visited = grabNumber(input, Point{x: symbol.x, y: symbol.y - 1}, visited)
			sum += number
		}
	}

	return sum
}

func checkTopLeft(input []string, symbol Point) bool {
	if symbol.x > 0 && symbol.y > 0 {
		if unicode.IsDigit(rune(input[symbol.x-1][symbol.y-1])) {
			return true
		}
	}
	return false
}

func checkTop(input []string, symbol Point) bool {
	if symbol.x > 0 {
		if unicode.IsDigit(rune(input[symbol.x-1][symbol.y])) {
			return true
		}
	}
	return false
}

func checkTopRight(input []string, symbol Point) bool {
	if symbol.x > 0 && symbol.y < len(input[symbol.x]) {
		if unicode.IsDigit(rune(input[symbol.x-1][symbol.y+1])) {
			return true
		}
	}
	return false
}

func checkRight(input []string, symbol Point) bool {
	if symbol.y < len(input[symbol.x]) {
		if unicode.IsDigit(rune(input[symbol.x][symbol.y+1])) {
			return true
		}
	}
	return false
}

func checkBottomRight(input []string, symbol Point) bool {
	if symbol.x < len(input) && symbol.y < len(input[symbol.x]) {
		if unicode.IsDigit(rune(input[symbol.x+1][symbol.y+1])) {
			return true
		}
	}
	return false
}

func checkBottom(input []string, symbol Point) bool {
	if symbol.x < len(input) {
		if unicode.IsDigit(rune(input[symbol.x+1][symbol.y])) {
			return true
		}
	}
	return false
}

func checkBottomLeft(input []string, symbol Point) bool {
	if symbol.x < len(input) && symbol.y > 0 {
		if unicode.IsDigit(rune(input[symbol.x+1][symbol.y-1])) {
			return true
		}
	}
	return false
}

func checkLeft(input []string, symbol Point) bool {
	if symbol.y > 0 {
		if unicode.IsDigit(rune(input[symbol.x][symbol.y-1])) {
			return true
		}
	}
	return false
}

func checkVisited(visited []Point, point Point) bool {
	for i := 0; i < len(visited); i++ {
		if visited[i].x == point.x && visited[i].y == point.y {
			return true
		}
	}
	return false
}

func grabNumber(input []string, startPoint Point, visited []Point) (int, []Point) {
	visited = append(visited, startPoint)
	numToReturnAsString := string(input[startPoint.x][startPoint.y])

	// Check left
	i := startPoint.y - 1
	for i >= 0 {
		if unicode.IsDigit(rune(input[startPoint.x][i])) {
			numToReturnAsString = string(input[startPoint.x][i]) + numToReturnAsString
			visited = append(visited, Point{x: startPoint.x, y: i})
		} else {
			break
		}
		i--
	}

	// Check Right
	i = startPoint.y + 1
	for i < len(input[startPoint.x]) {
		if unicode.IsDigit(rune(input[startPoint.x][i])) {
			numToReturnAsString += string(input[startPoint.x][i])
			visited = append(visited, Point{x: startPoint.x, y: i})
		} else {
			break
		}
		i++
	}

	numToReturn, err := strconv.Atoi(numToReturnAsString)
	util.Check(err)
	return numToReturn, visited
}
