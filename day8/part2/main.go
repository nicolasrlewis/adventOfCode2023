package main

import (
	"adventOfCode2023/util"
	"fmt"
	"strings"
)

type Node struct {
	label string
	left  *Node
	right *Node
	index int
}

// day8 Part 2
func main() {
	input, err := util.GetPart2InputFileLines("./day8")
	//input, err := util.GetExample2InputFileLines("./day8")
	util.Check(err)

	instructions := input[0]

	nodes := buildNodes(input[1:])

	nodeIndexes := findEndingANodes(nodes)

	var minSteps []int
	for i := 0; i < len(nodeIndexes); i++ {
		minSteps = append(minSteps, findSteps(nodes, nodeIndexes[i], instructions))
	}

	fmt.Printf("Minimum steps for part2 is : %d, use LCM for answer.\n", minSteps)
}

func buildNodes(input []string) []*Node {
	var nodes []*Node
	for i := 0; i < len(input); i++ {
		splitInput := strings.Split(input[i], " = ")

		label := splitInput[0]

		splitInput = strings.Split(splitInput[1], "(")
		splitInput = strings.Split(splitInput[1], ")")
		splitInput = strings.Split(splitInput[0], ", ")

		leftLabel := splitInput[0]
		rightLabel := splitInput[1]

		var node *Node
		for j := 0; j < len(nodes); j++ {
			if nodes[j].label == label {
				node = nodes[j]
				break
			}
		}
		if node == nil {
			node = new(Node)
			node.label = label
			nodes = append(nodes, node)
			node.index = len(nodes) - 1
		}

		var leftNode *Node
		if node.label != leftLabel {
			for j := 0; j < len(nodes); j++ {
				if nodes[j].label == leftLabel {
					leftNode = nodes[j]
					break
				}
			}
			if leftNode == nil {
				leftNode = new(Node)
				leftNode.label = leftLabel
				nodes = append(nodes, leftNode)
				leftNode.index = len(nodes) - 1
			}
		}
		var rightNode *Node
		if node.label != rightLabel {
			for j := 0; j < len(nodes); j++ {
				if nodes[j].label == rightLabel {
					rightNode = nodes[j]
					break
				}
			}
			if rightNode == nil {
				rightNode = new(Node)
				rightNode.label = rightLabel
				nodes = append(nodes, rightNode)
				rightNode.index = len(nodes) - 1
			}
		} else {
			rightNode = node
		}

		node.left = leftNode
		node.right = rightNode
	}

	return nodes
}

func findAAANodeIndex(nodes []*Node) int {
	for i := 0; i < len(nodes); i++ {
		if nodes[i].label == "AAA" {
			return i
		}
	}

	return -1
}

func findEndingANodes(nodes []*Node) []int {
	var endingANodes []int
	for i := 0; i < len(nodes); i++ {
		if nodes[i].label[len(nodes[i].label)-1:] == "A" {
			endingANodes = append(endingANodes, i)
		}
	}
	return endingANodes
}

func findSteps(nodes []*Node, nodeIndex int, instructions string) int {
	i := 0
	numOfSteps := 0
	for {
		if nodes[nodeIndex].label[len(nodes[nodeIndex].label)-1:] == "Z" {
			break
		}
		if string(instructions[i]) == "L" {
			nodeIndex = nodes[nodeIndex].left.index
			numOfSteps++
		} else {
			nodeIndex = nodes[nodeIndex].right.index
			numOfSteps++
		}

		i = (i + 1) % len(instructions)
	}

	return numOfSteps
}
