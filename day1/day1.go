package day1

import (
	"fmt"
	"os"

	"github.com/mike1808/adventofcode2020/util"
)

const TargetSum = 2020

func Day1() {
	nums, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Println("The answer for Day 1 part 1 = ", Sum2To2020(nums, TargetSum))
	fmt.Println("The answer for Day 1 part 2 = ", Sum3To2020(nums))
}

// time = O(n), memory = O(n)
func Sum2To2020(nums []int, target int) int {
	numsMap := map[int]struct{}{}

	for _, num := range nums {
		need := target - num
		if _, ok := numsMap[need]; ok {
			return need * num
		}
		numsMap[num] = struct{}{}
	}

	return 0
}

// time = O(n^2), memory = O(n)
func Sum3To2020(nums []int) int {
	seen := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = struct{}{}
			if x := Sum2To2020(nums, TargetSum-num); x != 0 {
				return x * num
			}
		}
	}

	return 0
}

func parseInput() ([]int, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return util.ReadIntSlice(f)
}
