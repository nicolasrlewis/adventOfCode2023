package main

import (
	"adventOfCode2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SourceToDestination struct {
	destinationStart int
	sourceStart      int
	rangeLen         int
}

// day5 Part 2
func main() {
	input, err := util.GetPart2InputFileLines("./day5")
	//input, err := util.GetExample2InputFileLines("./day5")
	util.Check(err)

	//seeds := getSeeds(input[0])

	var seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc []string
	index := 3

	seedToSoil = buildMappingList(input[index:])
	index += len(seedToSoil) + 2

	soilToFert = buildMappingList(input[index:])
	index += len(soilToFert) + 2

	fertToWater = buildMappingList(input[index:])
	index += len(fertToWater) + 2

	waterToLight = buildMappingList(input[index:])
	index += len(waterToLight) + 2

	lightToTemp = buildMappingList(input[index:])
	index += len(lightToTemp) + 2

	tempToHum = buildMappingList(input[index:])
	index += len(tempToHum) + 2

	humToLoc = buildMappingList(input[index:])
	index += len(humToLoc) + 2

	var seedToSoilSTD, soilToFertSTD, fertToWaterSTD, waterToLightSTD, lightToTempSTD, tempToHumSTD, humToLocSTD []SourceToDestination
	seedToSoilSTD = buildSourceToDestinationList(seedToSoil)
	soilToFertSTD = buildSourceToDestinationList(soilToFert)
	fertToWaterSTD = buildSourceToDestinationList(fertToWater)
	waterToLightSTD = buildSourceToDestinationList(waterToLight)
	lightToTempSTD = buildSourceToDestinationList(lightToTemp)
	tempToHumSTD = buildSourceToDestinationList(tempToHum)
	humToLocSTD = buildSourceToDestinationList(humToLoc)

	minLocation := math.MaxInt

	splitInput := strings.Split(input[0], "seeds: ")
	splitInput = strings.Split(splitInput[1], " ")

	counter := 1
	for i := 0; i < len(splitInput); i = i + 2 {
		start, _ := strconv.Atoi(splitInput[i])
		rangeLen, _ := strconv.Atoi(splitInput[i+1])

		fmt.Printf("Processing seed range #%d\n", counter)
		for k := start; k < start+rangeLen; k++ {
			location := processSeed(k, seedToSoilSTD, soilToFertSTD, fertToWaterSTD, waterToLightSTD, lightToTempSTD, tempToHumSTD, humToLocSTD)

			if location < minLocation {
				minLocation = location
			}
		}
		fmt.Printf("Finished processing seed range #%d\n", counter)
		counter++
	}

	fmt.Printf("The answer to part2 is: %d\n", minLocation)
}

func buildMappingList(input []string) []string {
	var mappingList []string

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			break
		}

		mappingList = append(mappingList, input[i])
	}

	return mappingList
}

func buildSourceToDestinationList(input []string) []SourceToDestination {
	var sourceToDestList []SourceToDestination

	for i := 0; i < len(input); i++ {
		var sourceToDest SourceToDestination
		splitInput := strings.Split(input[i], " ")

		sourceToDest.destinationStart, _ = strconv.Atoi(splitInput[0])
		sourceToDest.sourceStart, _ = strconv.Atoi(splitInput[1])
		sourceToDest.rangeLen, _ = strconv.Atoi(splitInput[2])

		sourceToDestList = append(sourceToDestList, sourceToDest)
	}

	return sourceToDestList
}

func processSeed(seed int, seedToSoil []SourceToDestination, soilToFert []SourceToDestination, fertToWater []SourceToDestination, waterToLight []SourceToDestination, lightToTemp []SourceToDestination,
	tempToHum []SourceToDestination, humToLoc []SourceToDestination) int {

	id := processMapping(seed, seedToSoil)
	id = processMapping(id, soilToFert)
	id = processMapping(id, fertToWater)
	id = processMapping(id, waterToLight)
	id = processMapping(id, lightToTemp)
	id = processMapping(id, tempToHum)
	id = processMapping(id, humToLoc)

	return id
}

func processMapping(id int, sourceToDestArray []SourceToDestination) int {

	index := -1
	for i := 0; i < len(sourceToDestArray); i++ {
		sourceRange := sourceToDestArray[i].sourceStart + sourceToDestArray[i].rangeLen

		if sourceToDestArray[i].sourceStart <= id && sourceRange > id {
			index = i
			break
		}
	}

	if index == -1 {
		return id
	} else {
		sourceToDest := sourceToDestArray[index]
		difference := int(math.Abs(float64(sourceToDest.sourceStart - sourceToDest.destinationStart)))

		if sourceToDest.destinationStart > sourceToDest.sourceStart {
			return id + difference
		} else {
			return id - difference
		}
	}

}
