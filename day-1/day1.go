package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

// FindSumAndMultiply ...
func FindSumAndMultiply(expenseReport []int) int {
	return 3
}

// FindItemsInListThatSumTo2020 ...
func FindItemsInListThatSumTo2020(expenseReport []int) (int, int) {
	target := 2020
	sort.Ints(expenseReport)
	for i, v := range expenseReport {
		v2 := target - v
		if j := sort.SearchInts(expenseReport, v2); v2 == expenseReport[j] {
			a := expenseReport[i]
			b := expenseReport[j]
			return a, b
		}
	}
	return -1, -1
}
