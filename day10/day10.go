package day10

import (
	"fmt"
	"os"
	"sort"

	"github.com/mike1808/adventofcode2020/util"
)

func Day10() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 10 part 1 answer is %d\n", Part1(input))
	fmt.Printf("Day 10 part 2 answer is %d\n", Part2(input))
}

func Part1(input []int) int {
	sort.Ints(input)

	prev := 0
	diff1 := 0
	diff3 := 1

	for _, n := range input {
		switch n - prev {
		case 1:
			diff1++
		case 3:
			diff3++
		default:
			panic("wtf")
		}
		prev = n
	}

	return diff1 * diff3
}

func Part2(input []int) int64 {
	sort.Ints(input)

	memo := map[string]int64{}

	var helper func([]int, int) int64
	helper = func(arr []int, prev int) int64 {
		if len(arr) == 0 {
			return 1
		}

		if v, ok := memo[key(len(arr), prev)]; ok {
			return v
		}

		c := int64(0)

		for i, n := range arr {
			if n-prev <= 3 {
				c += helper(arr[i+1:], n)
			} else {
				break
			}
		}

		memo[key(len(arr), prev)] = c
		return c
	}

	return helper(input, 0)
}

func key(l, p int) string {
	return fmt.Sprintf("%d,%d", l, p)
}

func parseInput() ([]int, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return util.ReadIntSlice(f)
}
