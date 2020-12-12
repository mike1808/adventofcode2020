package day1

import (
	"testing"
)

func TestSum2To2020(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	got := Sum2To2020(input, TargetSum)

	if got != 1721*299 {
		t.Errorf("Sum2To2020(input) = %d; want 1721 * 299 = %d", got, 1721*299)
	}
}

func TestSum3To2020(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	got := Sum3To2020(input)

	if got != 979*366*675 {
		t.Errorf("Sum3To2020(input) = %d; want 979*366*675 = %d", got, 979*366*675)
	}
}
func TestDay1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced")
		}
	}()
	Day1()
}
