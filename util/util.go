package util

import (
	"bufio"
	"io"
	"strconv"
)

func Max(a ...int) int {
	m := a[0]
	for _, n := range a[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func Min(a ...int) int {
	m := a[0]
	for _, n := range a[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func ReadIntSlice(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		result = append(result, n)
	}
	return result, scanner.Err()
}
