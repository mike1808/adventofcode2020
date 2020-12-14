package day6

import (
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day6()
}

func TestSumOfGroupAllYesAnswers(t *testing.T) {
	input := strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b`)
	actual := SumOfGroupAllYesAnswers(input)

	if actual != 6 {
		t.Errorf("SumOfGroupYesAnswers() returned: %d, expected: %d", actual, 6)
	}
}

func TestSumOfGroupYesAnswers(t *testing.T) {
	input := strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b`)
	actual := SumOfGroupYesAnswers(input)

	if actual != 11 {
		t.Errorf("SumOfGroupYesAnswers() returned: %d, expected: %d", actual, 11)
	}
}
