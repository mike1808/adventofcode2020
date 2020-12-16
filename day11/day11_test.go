package day11

import (
	"strings"
	"testing"
)

func TestDay11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day11()
}

func TestPart1(t *testing.T) {
	input, _ := readInput(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))
	actual := Part1(input)

	if actual != 37 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 37)
	}
}

func TestPart2(t *testing.T) {
	input, _ := readInput(strings.NewReader(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`))
	actual := Part2(input)

	if actual != 26 {
		t.Errorf("Part2() returned: %d, expected: %d", actual, 26)
	}
}
