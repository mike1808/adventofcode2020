package day16

import (
	"strings"
	"testing"
)

func TestDay16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day16()
}

func TestPart1(t *testing.T) {
	input, _ := readInput(strings.NewReader(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`))
	actual := Part1(input)

	if actual != 71 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 71)
	}
}

func TestPart2(t *testing.T) {
	input, _ := readInput(strings.NewReader(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`))
	actual := Part2(input)

	if actual != 71 {
		t.Errorf("Part2() returned: %d, expected: %d", actual, 71)
	}
}
