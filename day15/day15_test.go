package day15

import (
	"testing"
)

func TestDay15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day15()
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 3, 6}, 436},
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}
	for _, test := range tests {
		actual := Part1(test.input, 2020)
		if actual != test.expected {
			t.Errorf("Part1(%v) returned: %d, expected: %d", test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 3, 6}, 175594},
		{[]int{1, 3, 2}, 2578},
		{[]int{2, 1, 3}, 3544142},
		{[]int{1, 2, 3}, 261214},
		{[]int{2, 3, 1}, 6895259},
		{[]int{3, 2, 1}, 18},
		{[]int{3, 1, 2}, 362},
	}
	for _, test := range tests {
		actual := Part1(test.input, 30000000)
		if actual != test.expected {
			t.Errorf("Part1(%v) returned: %d, expected: %d", test.input, actual, test.expected)
		}
	}
}
