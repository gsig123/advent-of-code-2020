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

func TestFindItemsInListThatSumToTarget(t *testing.T) {
	expenseReport := []int{1721, 979, 366, 299, 675, 1456}
	desiredAnswerA, desiredAnswerB := 299, 1721
	ansA, ansB := FindItemsInListThatSumToTarget(expenseReport, 2020)

	if ansA != desiredAnswerA {
		t.Errorf("First number not correct, got: %d, wanted: %d", ansA, desiredAnswerA)
	}

	if ansB != desiredAnswerB {
		t.Errorf("Second number not correct, got: %d, wanted: %d", ansB, desiredAnswerB)
	}
}

func TestFind3NumbersThatSumToTarget(t *testing.T) {
	expenseReport := []int{1721, 979, 366, 299, 675, 1456}
	desiredA, desiredB, desiredC := 979, 675, 366
	ansA, ansB, ansC := Find3NumbersThatSumToTarget(expenseReport, 2020)

	if ansA != desiredA {
		t.Errorf("First number not correct, got: %d, wanted: %d", ansA, desiredA)
	}

	if ansB != desiredB {
		t.Errorf("Second number not correct, got: %d, wanted: %d", ansB, desiredB)
	}

	if ansC != desiredC {
		t.Errorf("Third number not correct, got: %d, wanted: %d", ansC, desiredC)
	}
}
