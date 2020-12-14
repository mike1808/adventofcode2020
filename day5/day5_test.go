package day5

import (
	"testing"
)

func TestDay5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day5()
}

func TestFindSeatsMaxID(t *testing.T) {
	input := []string{
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	actual := FindSeatsMaxID(input)

	if actual != 820 {
		t.Errorf("FindSeatsMaxID() returned: %d, expected: %d", actual, 820)
	}
}
