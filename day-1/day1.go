package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Read in file
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	lines := strings.Split(txt, "\n")

	// Turn it into ints
	var expenseReport = []int{}
	for _, i := range lines {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		expenseReport = append(expenseReport, j)
	}

	res := FindSumAndMultiply(expenseReport)
	fmt.Println("Part 1: The sum is ", res)

	numA, numB, numC := Find3NumbersThatSumToTarget(expenseReport, 2020)
	productOfThree := numA * numB * numC

	fmt.Println("Part 2: The product is: ", productOfThree)

}

// FindSumAndMultiply ...
func FindSumAndMultiply(expenseReport []int) int {
	target := 2020
	numberA, numberB := FindItemsInListThatSumToTarget(expenseReport, target)
	return numberA * numberB
}

// FindItemsInListThatSumToTarget ...
// Finds a pair of two numbers that sum up to `target` in
// a list (`expenseReport`).
// Complexity: Time: O(n). Space: O(n)
func FindItemsInListThatSumToTarget(expenseReport []int, target int) (int, int) {
	m := make(map[int]int)
	for _, v := range expenseReport {
		temp := target - v
		if _, ok := m[temp]; ok {
			return v, temp
		}
		m[v] = 1
	}
	return -1, -1
}

// Find3NumbersThatSumToTarget ...
// Finds 3 numbers that sum up to `target` in
// a list (`list`).
// Complexity: Time: O(n^2). Space: O(n)
func Find3NumbersThatSumToTarget(list []int, target int) (int, int, int) {
	for i, val := range list {
		m := make(map[int]int)
		currSum := target - val
		for j := i + 1; j <= len(list)-1; j++ {
			key := currSum - list[j]
			if _, ok := m[key]; ok {
				return list[i], list[j], currSum - list[j]
			}
			m[list[j]] = 1
		}
	}
	return -1, -1, -1
}
