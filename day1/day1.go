package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	return readInts(f)
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
