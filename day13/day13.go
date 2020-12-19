package day13

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const Empty = -1

type Input struct {
	Timestamp int
	Buses     []int
}

func Day13() {
	input, err := parseInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Day 13 part 1 answer is %d\n", Part1(input))
	fmt.Printf("Day 13 part 2 answer is %d\n", Part2(input))
}

func Part1(input *Input) int {
	busID := -1
	busArrival := math.MaxInt32
	for _, bus := range input.Buses {
		if bus == Empty {
			continue
		}
		if input.Timestamp%bus == 0 {
			busID = bus
		}

		closest := bus * int(math.Ceil(float64(input.Timestamp)/float64(bus)))
		if closest < busArrival {
			busID = bus
			busArrival = closest
		}
	}

	return busID * (busArrival - input.Timestamp)
}

// not working
func Part2(input *Input) int64 {
	panic("not working")
	d := int64(1)
	x := int64(1)
	pass := false

	for ; !pass; x += d {
		pass = true

		for i, bus := range input.Buses {
			if bus == Empty {
				continue
			}
			if (x+int64(i))%int64(bus) == 0 {
				c := int64(bus - i)
				if d%c != 0 {
					d *= c
					// fmt.Printf("d = %d, bus = %d, i = %d, bus - i = %d\n", d, bus, i, bus-i)
				}
			} else {
				pass = false
			}
		}
	}

	return x
}

func parseInput() (*Input, error) {
	f, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readInput(f)
}

func readInput(r io.Reader) (*Input, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	timestamp, _ := strconv.Atoi(lines[0])
	busesStr := strings.Split(lines[1], ",")
	buses := []int{}
	for _, busStr := range busesStr {
		if busStr != "x" {
			id, _ := strconv.Atoi(busStr)
			buses = append(buses, id)
		} else {
			buses = append(buses, Empty)
		}
	}

	return &Input{
		Timestamp: timestamp,
		Buses:     buses,
	}, nil
}
