package day18

import (
	"strings"
	"testing"
)

func TestDay18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day18()
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, test := range tests {
		input, _ := readInput(strings.NewReader(test.input))
		actual := Part1(input)
		if actual != test.expected {
			t.Errorf("Part1(%q) returned: %d, expected: %d", test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 + 2 + 3 + 4", 10},
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}
	for _, test := range tests {
		input, _ := readInput(strings.NewReader(test.input))
		actual := Part2(input)
		if actual != test.expected {
			t.Errorf("Part2(%q) returned: %d, expected: %d", test.input, actual, test.expected)
		}
	}
}
