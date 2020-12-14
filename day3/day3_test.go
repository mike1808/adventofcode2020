package day3

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day3()
}

func TestCountDifferentPathTrees(t *testing.T) {
	input, width := exampleInput()
	out := CountDifferentPathTrees(input, width)

	if out != 336 {
		t.Errorf("CountTrees(input) returned %d, expected %d", out, 336)
	}
}

func TestCountTrees(t *testing.T) {
	input, width := exampleInput()
	out := CountTrees(input, width, 3, 1)

	if out != 7 {
		t.Errorf("CountTrees(input) returned %d, expected %d", out, 7)
	}
}

func TestReadInput(t *testing.T) {
	r := strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
	out, width, err := readInput(r)
	if err != nil {
		t.Errorf("readInput(r) didn't expect to error: %v", err)
	}
	expected, expectedWidth := exampleInput()

	if width != expectedWidth {
		t.Errorf("readInput(r) expected: %d, got: %d", expectedWidth, width)
	}
	for i, actual := range out {
		if !reflect.DeepEqual(actual, expected[i]) {
			t.Errorf("readInput(r) expected: %b, actual: %b", expected[i], actual)
		}
	}
}

func exampleInput() ([]uint32, int) {
	return []uint32{
		uint32(0b00110000000),
		uint32(0b10001000100),
		uint32(0b01000010010),
		uint32(0b00101000101),
		uint32(0b01000110010),
		uint32(0b00101100000),
		uint32(0b01010100001),
		uint32(0b01000000001),
		uint32(0b10110001000),
		uint32(0b10001100001),
		uint32(0b01001000101),
	}, 11
}
