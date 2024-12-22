package day1

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"aoc2024.com/util"
)

type Solver struct {
	list1 []int
	list2 []int
}

func New(lines []string) Solver {
	var list1 []int
	var list2 []int

	for _, line := range lines {
		parts := strings.Fields(line)
		val1, err := strconv.Atoi(parts[0])
		util.CheckWithFatalError(err, "String parsing failed with error %v")
		val2, err := strconv.Atoi(parts[1])
		util.CheckWithFatalError(err, "String parsing failed with error %v")
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	slices.Sort(list1)
	slices.Sort(list2)
	return Solver{list1, list2}
}

func (s Solver) Part1() int {
	var sum float64 = 0
	for index, item := range s.list1 {
		sum += math.Abs(float64(item - s.list2[index]))
	}
	return int(sum)
}

func (s Solver) Part2() int {
	idx1 := 0
	idx2 := 0
	sum := 0

	for idx1 < len(s.list1) && idx2 < len(s.list2) {
		val := s.list1[idx1]
		cnt := 0
		for s.list2[idx2] <= val {
			if s.list2[idx2] == val {
				cnt += 1
			}
			idx2++
		}
		sum += cnt * val
		idx1++
	}
	return sum
}
