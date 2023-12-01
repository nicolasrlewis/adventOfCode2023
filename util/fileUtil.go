package util

import (
	"bufio"
	"os"
)

const (
	InputDirectory    = "input"
	Example1InputPath = "./" + InputDirectory + "/" + "example1.txt"
	Example2InputPath = "./" + InputDirectory + "/" + "example2.txt"
	Part1InputPath    = "./" + InputDirectory + "/" + "part1.txt"
	Part2InputPath    = "./" + InputDirectory + "/" + "part2.txt"
)

func GetPart1InputFileAsString(dayPath string) (string, error) {
	return GetInputFileAsString(dayPath, Part1InputPath)
}

func GetPart1InputFileLines(dayPath string) ([]string, error) {
	return GetInputFileLines(dayPath, Part1InputPath)
}

func GetPart2InputFileAsString(dayPath string) (string, error) {
	return GetInputFileAsString(dayPath, Part2InputPath)
}

func GetPart2InputFileLines(dayPath string) ([]string, error) {
	return GetInputFileLines(dayPath, Part2InputPath)
}

func GetExample1InputFileAsString(dayPath string) (string, error) {
	return GetInputFileAsString(dayPath, Example1InputPath)
}

func GetExample1InputFileLines(dayPath string) ([]string, error) {
	return GetInputFileLines(dayPath, Example1InputPath)
}

func GetExample2InputFileAsString(dayPath string) (string, error) {
	return GetInputFileAsString(dayPath, Example2InputPath)
}

func GetExample2InputFileLines(dayPath string) ([]string, error) {
	return GetInputFileLines(dayPath, Example2InputPath)
}

func GetInputFileAsString(dayPath, filePath string) (string, error) {
	path := dayPath + "/" + filePath
	inputText, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(inputText), err
}

func GetInputFileLines(dayPath string, filePath string) ([]string, error) {
	path := dayPath + "/" + filePath
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
