package day12

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Coord [2]int

type Direction int

const (
	East Direction = iota
	South
	West
	North
	Left
	Right
	Forward
)

type Entry struct {
	distance  int
	direction Direction
}

func Day12() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 12 part 1 answer is %d\n", SailAndGetDistance(input))
	fmt.Printf("Day 12 part 2 answer is %d\n", SailAndGetDistanceWithWaypoint(input))
}

func SailAndGetDistance(input []*Entry) int {
	coord := Coord{0, 0}

	directionToDelta := [][2]int{
		[2]int{1, 0},  // East
		[2]int{0, -1}, // South
		[2]int{-1, 0}, // West
		[2]int{0, 1},  // North
	}

	d := East
	delta := directionToDelta[d]

	for _, entry := range input {
		switch entry.direction {
		case Forward:
			coord[0] += delta[0] * entry.distance
			coord[1] += delta[1] * entry.distance
		case North, South, East, West:
			dd := directionToDelta[entry.direction]
			coord[0] += dd[0] * entry.distance
			coord[1] += dd[1] * entry.distance
		case Left, Right:
			revs := entry.distance / 90
			var newD int
			if entry.direction == Left {
				newD = (int(d) - revs) % 4
			} else {
				newD = (int(d) + revs) % 4
			}

			if newD < 0 {
				newD += 4
			}

			d = Direction(newD)
			delta = directionToDelta[newD]
		}
	}

	return calculateManhattanDistance(coord)
}

func SailAndGetDistanceWithWaypoint(input []*Entry) int {
	coord := Coord{0, 0}
	waypoint := Coord{10, 1}

	directionToDelta := [][2]int{
		[2]int{1, 0},  // East
		[2]int{0, -1}, // South
		[2]int{-1, 0}, // West
		[2]int{0, 1},  // North
	}

	// Rotation matrix for counter-clockwise (left) direction
	rotationMatrix := [][2][2]int{
		[2][2]int{ // 90
			[2]int{0, -1},
			[2]int{1, 0},
		},
		[2][2]int{ // 180
			[2]int{-1, 0},
			[2]int{0, -1},
		},
		[2][2]int{ // 270
			[2]int{0, 1},
			[2]int{-1, 0},
		},
	}

	for _, entry := range input {
		switch entry.direction {
		case Forward:
			coord[0] += waypoint[0] * entry.distance
			coord[1] += waypoint[1] * entry.distance
		case North, South, East, West:
			dd := directionToDelta[entry.direction]
			waypoint[0] += dd[0] * entry.distance
			waypoint[1] += dd[1] * entry.distance
		case Left, Right:
			// Counter clockwise rotation degress
			deg := entry.distance

			// If the direction is clockwise convert to counter clockwise
			if entry.direction == Right {
				deg = 360 - deg
			}

			m := rotationMatrix[deg/90-1]
			matmul(m, &waypoint)
		}
	}

	return calculateManhattanDistance(coord)
}

func matmul(matrix [2][2]int, vector *Coord) {
	v := *vector
	(*vector)[0] = v[0]*matrix[0][0] + v[1]*matrix[0][1]
	(*vector)[1] = v[0]*matrix[1][0] + v[1]*matrix[1][1]
}

func calculateManhattanDistance(coord Coord) int {
	return abs(coord[0]) + abs(coord[1])
}

func parseInput() ([]*Entry, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	return readInput(f)
}

func readInput(r io.Reader) ([]*Entry, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []*Entry
	mapping := map[byte]Direction{
		'N': North,
		'S': South,
		'E': East,
		'W': West,
		'L': Left,
		'R': Right,
		'F': Forward,
	}

	for scanner.Scan() {
		t := scanner.Text()
		d, err := strconv.Atoi(t[1:])
		if err != nil {
			return result, err
		}
		result = append(result, &Entry{
			direction: mapping[t[0]],
			distance:  d,
		})
	}
	return result, scanner.Err()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
