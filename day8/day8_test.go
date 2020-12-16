package day8

import (
	"strings"
	"testing"
)

func TestDay8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day8()
}

func TestGetAccBeforeLoop(t *testing.T) {
	input, _ := readInput(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	actual, _ := GetAccBeforeLoop(input)

	if actual != 5 {
		t.Errorf("GetAccBeforeLoop() returned: %d, expected: %d", actual, 5)
	}
}

func TestFixCodeAndGetAcc(t *testing.T) {
	input, _ := readInput(strings.NewReader(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	actual := FixCodeAndGetAcc(input)

	if actual != 8 {
		t.Errorf("GetAccBeforeLoop() returned: %d, expected: %d", actual, 8)
	}
}
