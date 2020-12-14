package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Day3() {
	input, slopeWidth, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 3 part 1 answer is %d\n", CountTrees(input, slopeWidth, 3, 1))
	fmt.Printf("Day 3 part 1 answer is %d\n", CountDifferentPathTrees(input, slopeWidth))
}

func CountDifferentPathTrees(input []uint32, slopeWidth int) int {
	paths := [][2]uint{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := 1

	for _, path := range paths {
		res *= CountTrees(input, slopeWidth, path[0], path[1])
	}

	return res
}

func CountTrees(input []uint32, slopeWidth int, right, bottom uint) int {
	startShift := uint(slopeWidth - 1)
	shift := startShift
	mask := uint32(1) << shift

	count := 0

	for i := 0; i < len(input); i += int(bottom) {
		row := input[i]

		count += int(row&mask) >> shift
		if shift < right {
			shift = startShift - (right - 1 - shift)
			mask = uint32(1) << shift
		} else {
			shift -= right
			mask = mask >> right
		}
	}
	return count
}

func parseInput() ([]uint32, int, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, 0, err
	}
	return readInput(f)
}

func readInput(r io.Reader) ([]uint32, int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []uint32
	var width int
	for scanner.Scan() {
		x := scanner.Text()
		width = len(x)
		row := parseRow(x)
		result = append(result, row)
	}
	return result, width, scanner.Err()
}

// Tree is 1, square is 0
func parseRow(s string) uint32 {
	row := uint32(0)
	for _, r := range s {
		switch r {
		case '#':
			row = (row << 1) | 1
		case '.':
			row = row << 1
		}
	}

	return row
}
