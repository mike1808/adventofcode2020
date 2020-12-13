package day2

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountValidPasswords1(t *testing.T) {
	out := CountValidPasswords1(exampleInput())

	if out != 2 {
		t.Errorf("CountValidPasswords1(input) should return %d, but returned %d", 2, out)
	}
}

func TestCountValidPasswords2(t *testing.T) {
	out := CountValidPasswords2(exampleInput())

	if out != 1 {
		t.Errorf("CountValidPasswords2(input) should return %d, but returned %d", 1, out)
	}
}

func TestDay2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day2()
}

func TestReadInput(t *testing.T) {
	r := strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`)
	out, err := readInput(r)
	if err != nil {
		t.Errorf("readInput(r) didn't expect to error: %v", err)
	}
	if len(out) != 3 {
		t.Errorf("readInput(r) didn't processed all lines, expected: %d, actual: %d", 3, len(out))
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
			pass: "abcde",
			policy: Policy{
				char: 'a',
				lo:   1,
				hi:   3,
			},
		},
		&Entry{
			pass: "cdefg",
			policy: Policy{
				char: 'b',
				lo:   1,
				hi:   3,
			},
		},
		&Entry{
			pass: "ccccccccc",
			policy: Policy{
				char: 'c',
				lo:   2,
				hi:   9,
			},
		},
	}

}
