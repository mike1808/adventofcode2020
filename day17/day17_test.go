package day17

import (
	"strings"
	"testing"
)

func TestDay17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day17()
}

func TestPart1(t *testing.T) {
	input, _ := readInput(strings.NewReader(`.#.
..#
###
`))
	actual := Part1(input)

	if actual != 112 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 112)
	}
}

func TestPart2(t *testing.T) {
	input, _ := readInput(strings.NewReader(`.#.
..#
###
`))
	actual := Part2(input)

	if actual != 848 {
		t.Errorf("Part2() returned: %d, expected: %d", actual, 848)
	}
}
