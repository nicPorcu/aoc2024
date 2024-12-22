package day2

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2024.com/util"
)

type Solver struct {
	all [][]int
}

func New(lines []string) Solver {
	var all [][]int
	for _, line := range lines {
		var sequence []int
		parts := strings.Fields(line)

		for _, item := range parts {
			new_item, err := strconv.Atoi(item)
			util.CheckWithFatalError(err, "String parsing resulted in error %v")
			sequence = append(sequence, new_item)
		}
		all = append(all, sequence)
	}
	return Solver{all: all}
}

func (s Solver) Part1() int {
	total := 0
	for _, line := range s.all {
		decreasing := line[1] < line[0]

		for i, val := range line {
			if i >= len(line)-1 {
				total += 1
				break
			}
			if !check(val, line[i+1], decreasing) {
				break
			}
		}
	}
	return total
}

func check(val1 int, val2 int, decreasing bool) bool {
	diff := val2 - val1
	decreasingMatch := (diff < 0) == decreasing
	return util.Abs(diff) >= 1 && util.Abs(diff) <= 3 && decreasingMatch
}

func checkAbs(val int, decreasing bool) bool {
	return util.Abs(val) <= 3 && util.Abs(val) >= 1 && (val < 0) == decreasing
}

func (s Solver) Part2() int {
	total := 0
	for _, line := range s.all {
		decreasing := line[len(line)-1] < line[0]
		var diff []int
		for i, val := range line {
			if i == len(line)-1 {
				break
			}
			diff = append(diff, line[i+1]-val)
		}
		var bad []int
		for i, val := range diff {
			if !checkAbs(val, decreasing) {
				bad = append(bad, i)
			}
		}

		if len(bad) == 2 {
			if bad[1]-bad[0] == 1 {
				if checkAbs(diff[bad[0]]+diff[bad[1]], decreasing) {
					total++
				}
			}
		} else if len(bad) == 1 {
			if (bad[0] == 0) || (bad[0] == len(diff)-1) {
				total++
			} else if checkAbs(diff[bad[0]]+diff[bad[0]+1], decreasing) || checkAbs(diff[bad[0]]+diff[bad[0]-1], decreasing) {
				total++
			} else {
				fmt.Println(line)
			}
			continue
		} else if len(bad) == 0 {
			total++
		}
	}
	return total
}
