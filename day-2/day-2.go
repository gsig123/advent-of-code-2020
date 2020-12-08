package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, dataLoadErr := ioutil.ReadFile("data.txt")
	if dataLoadErr != nil {
		panic(dataLoadErr)
	}

	dataString := string(data)
	lines := strings.Split(dataString, "\n")

	partOneNumberOfValidPasswords := 0
	partTwoNumberOfValidPasswords := 0

	for _, line := range lines {
		x := strings.Split(line, " ")
		numbers := strings.Split(x[0], "-")
		numA, numAErr := strconv.Atoi(numbers[0])
		if numAErr != nil {
			panic(numAErr)
		}
		numB, numBErr := strconv.Atoi(numbers[1])
		if numBErr != nil {
			panic(numBErr)
		}

		char := strings.Split(x[1], "")[0]
		password := x[2]

		if PartOnePasswordValidator(password, char, numB, numA) {
			partOneNumberOfValidPasswords++
		}

		if PartTwoPasswordValidator(password, char, numB, numA) {
			partTwoNumberOfValidPasswords++
		}

	}
	fmt.Println("Part One | Number of valid passwords: ", partOneNumberOfValidPasswords)
	fmt.Println("Part Two | Number of valid passwords: ", partTwoNumberOfValidPasswords)

}

// PartOnePasswordValidator ...
func PartOnePasswordValidator(password string, char string, maxTimes int, minTimes int) bool {
	passwordArray := strings.Split(password, "")
	passwordArrayWithoutLetter := filter(passwordArray, func(passwordChar string) bool {
		return passwordChar == char
	})
	length := len(passwordArrayWithoutLetter)
	if length >= minTimes && length <= maxTimes {
		return true
	}
	return false
}

// PartTwoPasswordValidator ...
func PartTwoPasswordValidator(password string, char string, idxA int, idxB int) bool {
	passwordArray := strings.Split(password, "")
	charAtIdxA := passwordArray[idxA-1] == char
	charAtIdxB := passwordArray[idxB-1] == char
	if charAtIdxA && !charAtIdxB || !charAtIdxA && charAtIdxB {
		return true
	}
	return false
}

// --- Utils ---
func filter(arr []string, cond func(string) bool) []string {
	result := []string{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
