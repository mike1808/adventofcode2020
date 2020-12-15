package day7

import (
	"strings"
	"testing"
)

func TestDay7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code paniced: %v", r)
		}
	}()
	Day7()
}

func TestCountBagColorsToContain(t *testing.T) {
	input, _ := readInput(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`))
	actual := CountBagColorsToContain(input, targetBagColor)

	if actual != 4 {
		t.Errorf("CountBagColorsToContain() returned: %d, expected: %d", actual, 4)
	}
}

func TestCountBagContained(t *testing.T) {
	input, _ := readInput(strings.NewReader(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`))
	actual := CountBagContained(input, targetBagColor)

	if actual != 32 {
		t.Errorf("CountBagContained() returned: %d, expected: %d", actual, 32)
	}

	input, _ = readInput(strings.NewReader(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`))
	actual = CountBagContained(input, targetBagColor)

	if actual != 126 {
		t.Errorf("CountBagContained() returned: %d, expected: %d", actual, 126)
	}
}
