package day12

import (
	"reflect"
	"strings"
	"testing"
)

func TestDay12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day12()
}

func TestSailAndGetDistance(t *testing.T) {
	out := SailAndGetDistance(exampleInput())

	if out != 25 {
		t.Errorf("SailAndGetDistance() returned %d, expected %d", out, 25)
	}
}

func TestSailAndGetDistanceWithWaypoint(t *testing.T) {
	out := SailAndGetDistanceWithWaypoint(exampleInput())

	if out != 286 {
		t.Errorf("SailAndGetDistance() returned %d, expected %d", out, 286)
	}
}

func TestCalculateManhattanDistance(t *testing.T) {
	cases := []struct {
		coord    Coord
		expected int
	}{
		{Coord{0, 0}, 0},
		{Coord{2, 0}, 2},
		{Coord{-5, 0}, 5},
		{Coord{5, 5}, 10},
		{Coord{-5, 5}, 10},
		{Coord{2, -5}, 7},
	}

	for _, c := range cases {
		out := calculateManhattanDistance(c.coord)
		if out != c.expected {
			t.Errorf("calculateManhattandistance([0,0]) got %d, expected %d", out, c.expected)
		}
	}
}

func TestReadInput(t *testing.T) {
	r := strings.NewReader(`F10
N3
F7
R90
F11`)
	out, err := readInput(r)

	if err != nil {
		t.Errorf("readInput(r) didn't expect to error: %v", err)
	}
	expected := exampleInput()
	for i, actual := range out {
		if !reflect.DeepEqual(actual, expected[i]) {
			t.Errorf("readInput(r) expected: %+#v, actual: %+#v", *expected[i], *actual)
		}
	}
}

func exampleInput() []*Entry {
	return []*Entry{
		&Entry{
			direction: Forward,
			distance:  10,
		},
		&Entry{
			direction: North,
			distance:  3,
		},
		&Entry{
			direction: Forward,
			distance:  7,
		},
		&Entry{
			direction: Right,
			distance:  90,
		},
		&Entry{
			direction: Forward,
			distance:  11,
		},
	}

}
