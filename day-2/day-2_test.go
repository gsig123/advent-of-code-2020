package main

import (
	"fmt"
	"testing"
)

func TestPartOnePasswordValidator(t *testing.T) {
	var tests = []struct {
		min, max int
		letter   string
		sequence string
		want     bool
	}{
		{1, 3, "a", "abcde", true},
		{1, 3, "b", "cdefg", false},
		{2, 9, "c", "ccccccccc", true},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("%d", idx)
		t.Run(testname, func(t *testing.T) {
			ans := PartOnePasswordValidator(tt.sequence, tt.letter, tt.max, tt.min)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestPartTwoPasswordValidator(t *testing.T) {
	var tests = []struct {
		min, max int
		letter   string
		sequence string
		want     bool
	}{
		{1, 3, "a", "abcde", true},
		{1, 3, "b", "cdefg", false},
		{2, 9, "c", "ccccccccc", false},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("%d", idx)
		t.Run(testname, func(t *testing.T) {
			ans := PartTwoPasswordValidator(tt.sequence, tt.letter, tt.max, tt.min)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
