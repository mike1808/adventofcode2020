package day13

import (
	"strings"
	"testing"
)

func TestDay13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day13()
}

func TestPart1(t *testing.T) {
	input, _ := readInput(strings.NewReader(`939
7,13,x,x,59,x,31,19`))
	actual := Part1(input)

	if actual != 295 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 295)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			input:    "939\n7,13,x,x,59,x,31,19",
			expected: 1068781,
		},
		{
			input:    "939\n17,x,13,19",
			expected: 3417,
		},
		{
			input:    "939\n67,7,59,61",
			expected: 754018,
		},
		{
			input:    "939\n67,x,7,59,61",
			expected: 779210,
		},
		{
			input:    "939\n67,7,x,59,61",
			expected: 1261476,
		},
		{
			input:    "939\n1789,37,47,1889",
			expected: 1202161486,
		},
	}

	for _, test := range tests {
		input, _ := readInput(strings.NewReader(test.input))
		actual := Part2(input)
		if actual != test.expected {
			t.Errorf("Par2(%v) returned %d, expected %d", input.Buses, actual, test.expected)
		}
	}
}
