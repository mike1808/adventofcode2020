package day5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func Day5() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 5 part 1 answer is %d\n", FindSeatsMaxID(input))
	fmt.Printf("Day 5 part 2 answer is %d\n", FindMySeatID(input))
}

func FindMySeatID(input []string) int {
	ids := []int{}

	for _, pass := range input {
		row, col := findSeat(pass)
		id := getID(row, col)
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for i, n := range ids[1:] {
		if n != ids[i]+1 {
			return n - 1
		}
	}

	return -1
}

func FindSeatsMaxID(input []string) int {
	maxID := 0

	for _, pass := range input {
		row, col := findSeat(pass)
		id := getID(row, col)
		maxID = max(maxID, id)
	}

	return maxID
}

const MaxRow = 127
const MaxCol = 7

func findSeat(pass string) (int, int) {
	lo, hi := 0, MaxRow
	row, col := 0, 0

	for _, char := range pass {
		if lo == hi {
			row = lo
			lo, hi = 0, MaxCol
		}

		mid := lo + (hi-lo)/2
		switch char {
		case 'F', 'L':
			hi = mid
		case 'B', 'R':
			lo = mid + 1
		}
	}

	col = lo

	return row, col
}

func getID(row, col int) int {
	return row*8 + col
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput() ([]string, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	return readInput(f)
}

func readInput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
