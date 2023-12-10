package main

import (
	"adventOfCode2023/util"
	"fmt"
)

type TileType string

const (
	VTPipe TileType = "|"
	HZPipe TileType = "-"
	NEPipe TileType = "L"
	NWPipe TileType = "J"
	SWPipe TileType = "7"
	SEPipe TileType = "F"
	Ground TileType = "."
	S      TileType = "S"
)

const (
	Up = iota
	Right
	Down
	Left
)

type Tile struct {
	tileType          TileType
	distanceFromStart int
	visited           bool
}

type Point struct {
	x, y int
}

type Grid struct {
	tiles [][]Tile
	start Point
}

type Node struct {
	point Point

	next *Node
	prev *Node
}

type Loop struct {
	start *Node
	size  int
}

// day10 Part 1
func main() {
	input, err := util.GetPart1InputFileLines("./day10")
	//input, err := util.GetExample1InputFileLines("./day10")
	util.Check(err)
	grid := buildGrid(input)

	loop := findLoop(grid)
	fmt.Printf("The answer for part 1 is: %d", loop.size/2)
}

func buildGrid(input []string) Grid {
	var grid Grid
	var tiles [][]Tile

	for i := 0; i < len(input); i++ {
		runes := []rune(input[i])

		row := make([]Tile, len(runes))
		for j := 0; j < len(runes); j++ {
			tileType := determineTileType(string(runes[j]))

			if tileType == S {
				grid.start = Point{
					x: i,
					y: j,
				}
			}

			row[j] = Tile{
				tileType:          tileType,
				distanceFromStart: -1,
			}
		}
		tiles = append(tiles, row)
	}

	grid.tiles = tiles
	return grid
}

func determineTileType(s string) TileType {
	switch s {
	case string(VTPipe):
		return VTPipe
	case string(HZPipe):
		return HZPipe
	case string(NEPipe):
		return NEPipe
	case string(NWPipe):
		return NWPipe
	case string(SWPipe):
		return SWPipe
	case string(SEPipe):
		return SEPipe
	case string(Ground):
		return Ground
	case string(S):
		return S
	default:
		panic("Invalid tileType")
	}
}

func findLoop(grid Grid) Loop {
	var loop Loop
	loop.start = &Node{
		point: grid.start,
		next:  nil,
		prev:  nil,
	}

	_, temp := progressFirst(grid, grid.start, &loop)
	loop.start.next = temp.next

	setLoopTileDistance(&grid, loop)

	return loop
}

func canMoveUp(grid Grid, point Point, validMoves []int) bool {
	if point.x == 0 {
		return false
	}
	if grid.tiles[point.x-1][point.y].visited {
		return false
	}
	if !contains(validMoves, Up) {
		return false
	}

	destination := grid.tiles[point.x-1][point.y].tileType

	switch destination {
	case VTPipe:
		return true
	case SWPipe:
		return true
	case SEPipe:
		return true
	case S:
		return true
	default:
		return false
	}
}

func canMoveRight(grid Grid, point Point, validMoves []int) bool {
	if point.y >= len(grid.tiles[0])-1 {
		return false
	}

	if grid.tiles[point.x][point.y+1].visited {
		return false
	}
	if !contains(validMoves, Right) {
		return false
	}

	destination := grid.tiles[point.x][point.y+1].tileType

	switch destination {
	case HZPipe:
		return true
	case NWPipe:
		return true
	case SWPipe:
		return true
	case S:
		return true
	default:
		return false
	}
}
func canMoveDown(grid Grid, point Point, validMoves []int) bool {
	if point.x >= len(grid.tiles)-1 {
		return false
	}
	if grid.tiles[point.x+1][point.y].visited {
		return false
	}
	if !contains(validMoves, Down) {
		return false
	}

	destination := grid.tiles[point.x+1][point.y].tileType

	switch destination {
	case VTPipe:
		return true
	case NEPipe:
		return true
	case NWPipe:
		return true
	case S:
		return true
	default:
		return false
	}
}
func canMoveLeft(grid Grid, point Point, validMoves []int) bool {
	if point.y <= 0 {
		return false
	}
	if grid.tiles[point.x][point.y-1].visited {
		return false
	}
	if !contains(validMoves, Left) {
		return false
	}

	destination := grid.tiles[point.x][point.y-1].tileType

	switch destination {
	case HZPipe:
		return true
	case NEPipe:
		return true
	case SEPipe:
		return true
	case S:
		return true
	default:
		return false
	}
}

func progressFirst(grid Grid, point Point, loop *Loop) (bool, *Node) {

	validMoves := buildValidMoves(grid.tiles[point.x][point.y].tileType)

	if canMoveUp(grid, point, validMoves) {

		foundStart, temp := progress(grid, Point{
			x: point.x - 1,
			y: point.y,
		}, loop)

		if foundStart {
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			return true, newNode
		}

	}
	if canMoveRight(grid, point, validMoves) {
		foundStart, temp := progress(grid, Point{
			x: point.x,
			y: point.y + 1,
		}, loop)

		if foundStart {
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			return true, newNode
		}
	}
	if canMoveDown(grid, point, validMoves) {
		foundStart, temp := progress(grid, Point{
			x: point.x + 1,
			y: point.y,
		}, loop)

		if foundStart {
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			return true, newNode
		}
	}
	if canMoveLeft(grid, point, validMoves) {
		foundStart, temp := progress(grid, Point{
			x: point.x,
			y: point.y - 1,
		}, loop)

		if foundStart {
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			return true, newNode
		}
	}
	return false, nil
}

func progress(grid Grid, point Point, loop *Loop) (bool, *Node) {
	grid.tiles[point.x][point.y].visited = true

	if grid.tiles[point.x][point.y].tileType == S {
		loop.size++
		return true, loop.start
	}

	validMoves := buildValidMoves(grid.tiles[point.x][point.y].tileType)

	if canMoveUp(grid, point, validMoves) {

		foundStart, temp := progress(grid, Point{
			x: point.x - 1,
			y: point.y,
		}, loop)

		if foundStart {
			if point.x == 60 && point.y == 74 {
				fmt.Printf("WHAT")
			}
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			loop.size++
			return true, newNode
		}

	}
	if canMoveRight(grid, point, validMoves) {
		foundStart, temp := progress(grid, Point{
			x: point.x,
			y: point.y + 1,
		}, loop)

		if foundStart {
			if point.x == 60 && point.y == 74 {
				fmt.Printf("WHAT")
			}
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			loop.size++
			return true, newNode
		}
	}
	if canMoveDown(grid, point, validMoves) {
		foundStart, temp := progress(grid, Point{
			x: point.x + 1,
			y: point.y,
		}, loop)

		if foundStart {
			if point.x == 60 && point.y == 74 {
				fmt.Printf("WHAT")
			}
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			loop.size++
			return true, newNode
		}
	}
	if canMoveLeft(grid, point, validMoves) {
		foundStart, temp := progress(grid, Point{
			x: point.x,
			y: point.y - 1,
		}, loop)

		if foundStart {
			if point.x == 60 && point.y == 74 {
				fmt.Printf("WHAT")
			}
			newNode := &Node{
				point: point,
				next:  temp,
				prev:  nil,
			}
			temp.prev = newNode
			loop.size++
			return true, newNode
		}
	}
	return false, nil
}

func setLoopTileDistance(grid *Grid, loop Loop) {
	counter := 0
	grid.tiles[loop.start.point.x][loop.start.point.y].distanceFromStart = counter

	flag := false
	for curNode := loop.start.next; curNode != loop.start; curNode = curNode.next {
		if counter >= loop.size/2 {
			flag = true
		}
		if flag {
			counter--
		} else {
			counter++
		}
		if curNode.point.x == 60 && curNode.point.y == 74 {
			fmt.Printf("WHAT")
		}
		grid.tiles[curNode.point.x][curNode.point.y].distanceFromStart = counter
	}
}

func buildValidMoves(tileType TileType) []int {
	switch tileType {
	case VTPipe:
		return []int{Up, Down}
	case HZPipe:
		return []int{Right, Left}
	case NEPipe:
		return []int{Up, Right}
	case NWPipe:
		return []int{Up, Left}
	case SWPipe:
		return []int{Down, Left}
	case SEPipe:
		return []int{Down, Right}
	case S:
		return []int{Up, Right, Down, Left}
	default:
		return []int{}
	}
}

func contains(array []int, val int) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}
