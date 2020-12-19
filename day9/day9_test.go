package day9

import (
	"strings"
	"testing"

	"github.com/mike1808/adventofcode2020/util"
)

func TestDay9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day9()
}

func TestFindFirstInvalidNumber(t *testing.T) {
	input, _ := util.ReadIntSlice(strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`))
	actual := FindFirstInvalidNumber(input, 5)

	if actual != 127 {
		t.Errorf("FindFirstInvalidNumber() returned: %d, expected: %d", actual, 127)
	}
}

func TestCrackTheCode(t *testing.T) {
	input, _ := util.ReadIntSlice(strings.NewReader(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`))
	actual := CrackTheCode(input, 5)

	if actual != 62 {
		t.Errorf("CrackTheCode() returned: %d, expected: %d", actual, 62)
	}
}
