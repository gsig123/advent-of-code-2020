package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var testDataPath = "test-data.txt"
var dataPath = "data.txt"

func main() {
	data, dataLoadErr := ioutil.ReadFile(dataPath)
	if dataLoadErr != nil {
		panic(dataLoadErr)
	}

	slope := parseInput(data)

	steps := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treesHit := []int{}

	for step := 0; step < len(steps); step++ {
		treesHit = append(treesHit, getNumberOfTreesHit(slope, steps[step][0], steps[step][1]))
	}

	multiplied := 1
	for tree := 0; tree < len(treesHit); tree++ {
		multiplied = multiplied * treesHit[tree]
	}

	fmt.Println("Part 1 | Number of trees: ", getNumberOfTreesHit(slope, 3, 1))
	fmt.Println("Part 2 | Multiplied trees hit: ", multiplied)
}

func calcXCoordinate(xCoordinate int, patternLength int) int {
	return int(math.Mod(float64(xCoordinate), float64(patternLength)))
}

func getNumberOfTreesHit(slope [][]string, xStep int, yStep int) int {
	xCoordinate := 0
	yCoordinate := 0
	patternLength := len(slope[0])
	slopeLength := len(slope)
	numberOfTrees := 0
	for yCoordinate < slopeLength {
		if !(xCoordinate == 0) {
			x := calcXCoordinate(xCoordinate, patternLength)
			if slope[yCoordinate][x] == "#" {
				numberOfTrees++
			}
		}
		xCoordinate += xStep
		yCoordinate += yStep
	}

	return numberOfTrees
}

func parseInput(data []byte) [][]string {
	dataString := string(data)
	lines := strings.Split(dataString, "\n")
	var outerArray [][]string
	for _, line := range lines {
		var innerArray []string
		lineChars := strings.Split(line, "")
		for _, char := range lineChars {
			innerArray = append(innerArray, char)
		}
		outerArray = append(outerArray, innerArray)
	}
	return outerArray
}
