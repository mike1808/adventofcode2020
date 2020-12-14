package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Day6() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf("Day 6 part 1 answer is %d\n", SumOfGroupYesAnswers(f))
	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 6 part 2 answer is %d\n", SumOfGroupAllYesAnswers(f))
}

func SumOfGroupAllYesAnswers(r io.Reader) int {
	sum := 0

	_ = processInput(r, func(group []string) {
		answers := map[rune]int{}
		count := 0

		for _, g := range group {
			for _, q := range g {
				answers[q]++

				if answers[q] == len(group) {
					count++
				}
			}
		}

		sum += count
	})

	return sum
}

func SumOfGroupYesAnswers(r io.Reader) int {
	sum := 0

	_ = processInput(r, func(group []string) {
		answers := map[rune]struct{}{}
		count := 0
		for _, g := range group {
			for _, q := range g {
				if _, ok := answers[q]; !ok {
					count++
					answers[q] = struct{}{}
				}
			}
		}
		sum += count
	})

	return sum
}

type groupProcessor func([]string)

func processInput(r io.Reader, callback groupProcessor) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var group []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			callback(group)
			group = nil
		} else {
			group = append(group, line)
		}
	}

	if group != nil {
		callback(group)
	}

	return scanner.Err()
}
