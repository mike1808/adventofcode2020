package day14

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestDay14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day14()
}

func TestPart1(t *testing.T) {
	input, _ := readInput(ioutil.NopCloser(strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)))
	actual := Part1(input)

	if actual != 165 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 165)
	}
}

func TestPart2(t *testing.T) {
	input, _ := readInput(ioutil.NopCloser(strings.NewReader(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)))
	actual := Part2(input)

	if actual != 208 {
		t.Errorf("Part1() returned: %d, expected: %d", actual, 208)
	}
}
