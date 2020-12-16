package day9

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Day9() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 9 part 1 answer is %d\n", FindFirstInvalidNumber(input, 25))
	fmt.Printf("Day 9 part 2 answer is %d\n", CrackTheCode(input, 25))
}

func FindFirstInvalidNumber(input []int, windowSize int) int {
	seen := map[int]int{}

	for i, targetSum := range input[windowSize:] {
		valid := false
		for _, n := range input[i : i+windowSize] {
			compliment := targetSum - n
			if seen[compliment] > 0 {
				valid = true
				break
			}
			seen[n]++
		}

		if !valid {
			return targetSum
		}

		seen[input[i]]--
		seen[targetSum]++
	}

	return -1
}

func CrackTheCode(input []int, windowSize int) int {
	targetSum := FindFirstInvalidNumber(input, windowSize)

	start := 0
	end := 1
	sum := input[start]

	for i, n := range input[1:] {
		sum += n
		for sum > targetSum && start < i {
			sum -= input[start]
			start++
		}
		if sum == targetSum {
			end = i + 1
			break
		}
	}

	return max(input[start:end]...) + min(input[start:end]...)
}

func max(a ...int) int {
	m := a[0]
	for _, n := range a[1:] {
		if n > m {
			m = n
		}
	}
	return m
}
func min(a ...int) int {
	m := a[0]
	for _, n := range a[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func parseInput() ([]int, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readInput(f)
}

func readInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		result = append(result, n)
	}
	return result, scanner.Err()
}
