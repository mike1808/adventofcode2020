package day15

import (
	"fmt"
)

func Day15() {
	input := []int{0, 1, 5, 10, 3, 12, 19}
	fmt.Printf("Day 15 part 1 answer is %d\n", Part1(input, 2020))
	fmt.Printf("Day 15 part 2 answer is %d\n", Part1(input, 30000000))
}

func Part1(input []int, untilTurn int) int {
	spoken := make([]int, untilTurn)
	turn := 1

	for _, input := range input[:len(input)-1] {
		spoken[input] = turn
		turn++
	}

	var prev int
	speak := input[len(input)-1]

	for ; turn <= untilTurn; turn++ {
		prev = speak

		if t := spoken[speak]; t != 0 {
			speak = turn - t
		} else {
			speak = 0
		}

		spoken[prev] = turn
	}

	return prev
}
