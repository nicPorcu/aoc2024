package day7

import (
	"strconv"
	"strings"

	"aoc2024.com/util"
)

type Solver struct {
	values  [][]int
	targets []int
}

func New(lines []string) Solver {
	var targets []int
	var values [][]int

	for _, line := range lines {
		vals := strings.Split(line, ":")
		if len(vals) != 2 {
			panic("Invalid length")
		}
		target, err := strconv.Atoi(vals[0])
		util.CheckWithFatalError(err, "Invalid Integer")
		targets = append(targets, target)
		var line_vals []int
		for _, val := range strings.Split(strings.Trim(vals[1], " "), " ") {
			val_int, err := strconv.Atoi(val)
			util.CheckWithFatalError(err, "Invalid Integer")
			line_vals = append(line_vals, val_int)
		}
		values = append(values, line_vals)
	}
	return Solver{
		targets: targets,
		values:  values,
	}
}

func (s Solver) Part1() int {
	cnt := 0
	numLines := len(s.targets)
	for i := 0; i < numLines; i++ {
		var dynamicMap []map[int]struct{}
		initialMap := make(map[int]struct{})
		initialMap[0] = struct{}{}
		dynamicMap = append(dynamicMap, initialMap)
		line := s.values[i]
		for j, val := range line {
			lineMap := make(map[int]struct{})
			for prev, _ := range dynamicMap[j] {
				lineMap[prev+val] = struct{}{}
				lineMap[prev*val] = struct{}{}
			}
			dynamicMap = append(dynamicMap, lineMap)
		}
		_, exists := dynamicMap[len(line)][s.targets[i]]
		if exists {
			cnt += s.targets[i]
		}
	}
	return cnt
}

func (s Solver) Part2() int {
	cnt := 0
	numLines := len(s.targets)
	for i := 0; i < numLines; i++ {
		var dynamicMap []map[int]struct{}
		initialMap := make(map[int]struct{})
		initialMap[0] = struct{}{}
		dynamicMap = append(dynamicMap, initialMap)
		line := s.values[i]
		for j, val := range line {
			lineMap := make(map[int]struct{})
			for prev, _ := range dynamicMap[j] {
				prevStr := strconv.Itoa(prev)
				valStr := strconv.Itoa(val)
				concat, err := strconv.Atoi(prevStr + valStr)
				util.CheckWithFatalError(err, "Invalid Integer")
				lineMap[concat] = struct{}{}
				lineMap[prev+val] = struct{}{}
				lineMap[prev*val] = struct{}{}
			}
			dynamicMap = append(dynamicMap, lineMap)
		}
		_, exists := dynamicMap[len(line)][s.targets[i]]
		if exists {
			cnt += s.targets[i]
		}
	}
	return cnt
}
