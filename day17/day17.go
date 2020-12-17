package day17

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	Inactive = '.'
	Active   = '#'
)

type Grid1D map[int]byte
type Grid2D map[int]Grid1D
type Grid3D map[int]Grid2D
type Grid4D map[int]Grid3D

func Day17() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 17 part 1 answer is %d\n", Part1(input))
	fmt.Printf("Day 17 part 2 answer is %d\n", Part2(input))
}

const Cycles = 6

func Part1(input [][]byte) int {
	activeCubes := 0
	grid3D := make3DGrid(input)

	yRange := [2]int{-1, len(input)}
	xRange := [2]int{-1, len(input[0])}
	zRange := [2]int{-1, 1}

	for cycle := 1; cycle <= Cycles; cycle++ {
		activeCubes = 0

		newGrid3D := Grid3D{}

		for z := zRange[0]; z <= zRange[1]; z++ {
			newGrid3D[z] = Grid2D{}
			for y := yRange[0]; y <= yRange[1]; y++ {
				newGrid3D[z][y] = Grid1D{}
				for x := xRange[0]; x <= xRange[1]; x++ {
					nearbyActive := countActiveNeighbours3D(x, y, z, grid3D, true)
					cube := grid3D[z][y][x]
					var newCube byte

					switch cube {
					case Active:
						if nearbyActive == 2 || nearbyActive == 3 {
							newCube = Active
						} else {
							newCube = Inactive
						}
					case Inactive, byte(0):
						if nearbyActive == 3 {
							newCube = Active
						} else {
							newCube = Inactive
						}
					}
					newGrid3D[z][y][x] = newCube
					if newCube == Active {
						activeCubes++
					}
				}
			}
		}

		grid3D = newGrid3D
		xRange[0]--
		xRange[1]++
		yRange[0]--
		yRange[1]++
		zRange[0]--
		zRange[1]++
	}

	return activeCubes
}

func Part2(input [][]byte) int {
	activeCubes := 0
	grid4D := make4DGrid(input)

	yRange := [2]int{-1, len(input)}
	xRange := [2]int{-1, len(input[0])}
	zRange := [2]int{-1, 1}
	wRange := [2]int{-1, 1}

	for cycle := 1; cycle <= Cycles; cycle++ {
		activeCubes = 0

		newGrid4D := Grid4D{}

		for w := wRange[0]; w <= wRange[1]; w++ {
			newGrid4D[w] = Grid3D{}
			for z := zRange[0]; z <= zRange[1]; z++ {
				newGrid4D[w][z] = Grid2D{}
				for y := yRange[0]; y <= yRange[1]; y++ {
					newGrid4D[w][z][y] = Grid1D{}
					for x := xRange[0]; x <= xRange[1]; x++ {
						nearbyActive := countActiveNeighbours4D(x, y, z, w, grid4D)
						cube := grid4D[w][z][y][x]

						var newCube byte
						switch cube {
						case Active:
							if nearbyActive == 2 || nearbyActive == 3 {
								newCube = Active
							} else {
								newCube = Inactive
							}
						case Inactive, byte(0):
							if nearbyActive == 3 {
								newCube = Active
							} else {
								newCube = Inactive
							}
						}

						newGrid4D[w][z][y][x] = newCube
						if newCube == Active {
							activeCubes++
						}
					}
				}
			}
		}

		grid4D = newGrid4D
		xRange[0]--
		xRange[1]++
		yRange[0]--
		yRange[1]++
		zRange[0]--
		zRange[1]++
		wRange[0]--
		wRange[1]++
	}

	return activeCubes
}

func make3DGrid(grid [][]byte) Grid3D {
	grid3D := Grid3D{}
	grid3D[0] = Grid2D{}

	for y, row := range grid {
		grid3D[0][y] = Grid1D{}
		for x, v := range row {
			grid3D[0][y][x] = v
		}
	}

	return grid3D
}

func make4DGrid(grid [][]byte) Grid4D {
	grid4D := Grid4D{}
	grid4D[0] = make3DGrid(grid)
	return grid4D
}

func countActiveNeighbours3D(x, y, z int, grid Grid3D, skipSelf bool) int {
	active := 0

	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if skipSelf && dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				if grid2D, ok := grid[z+dz]; ok {
					if grid1D, ok := grid2D[y+dy]; ok {
						if value, ok := grid1D[x+dx]; ok && value == Active {
							active++
						}
					}
				}
			}
		}
	}

	return active
}

func countActiveNeighbours4D(x, y, z, w int, grid Grid4D) int {
	active := 0

	for dw := -1; dw <= 1; dw++ {
		if grid3D, ok := grid[w+dw]; ok {
			active += countActiveNeighbours3D(x, y, z, grid3D, dw == 0)
		}
	}

	return active
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
