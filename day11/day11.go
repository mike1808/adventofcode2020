package day11

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	Floor    = '.'
	Empty    = 'L'
	Occupied = '#'
)

func Day11() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 11 part 1 answer is %d\n", Part1(input))
	fmt.Printf("Day 11 part 2 answer is %d\n", Part2(input))
}

func Part1(grid [][]byte) int {
	occupiedSeats := 0
	changed := true

	for changed {
		occupiedSeats = 0
		changed = false

		newGrid := make([][]byte, len(grid))

		for i, row := range grid {
			newGrid[i] = make([]byte, len(row))

			for j, space := range row {
				newGrid[i][j] = space

				switch space {
				case Empty:
					occupied := countAdjOccupiedSeats(i, j, grid, false)
					if occupied == 0 {
						newGrid[i][j] = Occupied
						occupiedSeats++
						changed = true
					}
				case Occupied:
					occupied := countAdjOccupiedSeats(i, j, grid, false)
					if occupied >= 4 {
						newGrid[i][j] = Empty
						changed = true
					} else {
						occupiedSeats++
					}
				}
			}
		}

		grid = newGrid
	}

	return occupiedSeats
}

func Part2(grid [][]byte) int {
	occupiedSeats := 0
	changed := true

	for changed {
		occupiedSeats = 0
		changed = false

		newGrid := make([][]byte, len(grid))

		for i, row := range grid {
			newGrid[i] = make([]byte, len(row))

			for j, space := range row {
				newGrid[i][j] = space

				switch space {
				case Empty:
					occupied := countAdjOccupiedSeats(i, j, grid, true)
					if occupied == 0 {
						newGrid[i][j] = Occupied
						occupiedSeats++
						changed = true
					}
				case Occupied:
					occupied := countAdjOccupiedSeats(i, j, grid, true)
					if occupied >= 5 {
						newGrid[i][j] = Empty
						changed = true
					} else {
						occupiedSeats++
					}
				}
			}
		}

		grid = newGrid
	}

	return occupiedSeats
}

var deltas = [][2]int{
	{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1},
}

func countAdjOccupiedSeats(i, j int, grid [][]byte, part2 bool) int {
	m, n := len(grid), len(grid[0])
	count := 0

	for _, delta := range deltas {
		for ii, jj := i+delta[0], j+delta[1]; ii >= 0 && ii < m && jj >= 0 && jj < n; ii, jj = ii+delta[0], jj+delta[1] {
			if grid[ii][jj] != Floor {
				if grid[ii][jj] == Occupied {
					count++
				}
				break
			}

			if !part2 {
				break
			}
		}
	}

	return count
}

func parseInput() ([][]byte, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readInput(f)
}

func readInput(r io.Reader) ([][]byte, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result [][]byte
	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}
	return result, scanner.Err()
}
