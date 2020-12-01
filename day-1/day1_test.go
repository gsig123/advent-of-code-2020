package main

import "testing"

func TestSumAndMultiply(t *testing.T) {
	expenseReport := []int{1721, 979, 366, 299, 675, 1456}
	desiredAnswer := 514579
	ans := FindSumAndMultiply(expenseReport)

	if ans != desiredAnswer {
		t.Errorf("TestSumAndMultiply gave wrong answer, got: %d, wanted: %d.", ans, desiredAnswer)
	}
}

func TestFindItemsInListThatSumTo2020(t *testing.T) {
	expenseReport := []int{1721, 979, 366, 299, 675, 1456}
	desiredAnswerA, desiredAnswerB := 299, 1721
	ansA, ansB := FindItemsInListThatSumTo2020(expenseReport)

	if ansA != desiredAnswerA {
		t.Errorf("First number not correct, got: %d, wanted: %d", ansA, desiredAnswerA)
	}

	if ansB != desiredAnswerB {
		t.Errorf("Second number not correct, got: %d, wanted: %d", ansB, desiredAnswerB)
	}
}
